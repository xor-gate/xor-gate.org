---
title: "A Roundcube addressbook plugin with static contacts"
date: 2026-09-04
tags:
  - roundcube
  - webmail
  - php
  - plugins
toc: true
---

Roundcube ships with an [`example_addressbook`](https://github.com/roundcube/roundcubemail/tree/master/plugins/example_addressbook) plugin which exposes a hardcoded list of contacts as an extra address source. It is a great starting point when you want to expose an internal directory (LDAP is not always the answer) inside the webmail. What the upstream example does not show is how to return *structured* fields, like a postal address with a `home` and `work` variant, or how to add fields which Roundcube does not know about at all.

Below is a small, complete plugin which does exactly that: a read-only addressbook with static contacts, each carrying a home and a work address plus two plugin defined fields (a text and a date field, prefixed with `vtvbb_`) which are shown in the contact GUI.

## The address coltype

An addressbook backend advertises the fields it supports through the `$coltypes` property of `rcube_addressbook`. The default is minimal:

```php
public $coltypes = [
    'name'      => ['limit' => 1],
    'firstname' => ['limit' => 1],
    'surname'   => ['limit' => 1],
    'email'     => ['limit' => 1],
];
```

Roundcube merges this with its own general definition of contact columns (`rcmail_action_contacts_index::$CONTACT_COLTYPES`), where `address` is defined as a *composite* column with the subtypes `home`, `work` and `other` and the child fields `street`, `locality`, `zipcode`, `region` and `country`:

```php
'address' => [
    'type'     => 'composite',
    'label'    => 'address',
    'subtypes' => ['home', 'work', 'other'],
    'category' => 'main',
    'childs'   => [
        'street'   => ['label' => 'street',   'size' => 40, 'maxlength' => 50],
        'locality' => ['label' => 'locality', 'size' => 28, 'maxlength' => 50],
        'zipcode'  => ['label' => 'zipcode',  'size' => 8,  'maxlength' => 15],
        'region'   => ['label' => 'region',   'size' => 12, 'maxlength' => 50],
        'country'  => ['label' => 'country',  'size' => 40, 'maxlength' => 50],
    ],
],
```

Two things follow from this:

* Columns *not* listed in the backend `$coltypes` are removed from the UI, so `address` must be added explicitly.
* Whatever you put in your own `$coltypes['address']` is merged on top of the definition above, which is how you restrict the subtypes to `home` and `work` only.

In a contact record the subtype is part of the array key, and because a subtype may occur more than once the value is a *list* of value arrays:

```php
'address:home' => [
    [
        'street'   => 'Dorpsstraat 1',
        'locality' => 'Nijmegen',
        'zipcode'  => '6511 AA',
        'region'   => 'Gelderland',
        'country'  => 'Netherlands',
    ],
],
```

Roundcube collects these with `rcube_addressbook::get_col_values('address', $record)`, which splits every `address:*` key on the colon and groups the values per subtype. So `address`, `address:home` and `address:work` all end up in the same field group.

## Plugin defined custom fields

The merge of `$coltypes` and `$CONTACT_COLTYPES` is not limited to the columns Roundcube knows about. Columns which only exist in the backend definition are simply appended:

```php
if (isset(self::$CONTACT_COLTYPES[$col])) {
    self::$CONTACT_COLTYPES[$col] = array_merge(self::$CONTACT_COLTYPES[$col], $colprop);
} else {
    self::$CONTACT_COLTYPES[$col] = $colprop;
}
```

That is the extension point for plugin specific data. I prefix such columns with the plugin domain, `vtvbb_`, so they cannot collide with a current or future Roundcube column:

* `vtvbb_membership_id` — a plain text field
* `vtvbb_member_since` — a date field

Because Roundcube has no defaults for these columns, the backend must supply the complete definition: `type` (`text`, `date`, `select`, `composite`, ...), the input `size`/`maxlength`, a `limit` (1 means a single value, no *add field* button), an optional `category` and, for anything that is not plain text, a `render_func` which formats the stored value for display. For dates that is `rcmail_action_contacts_index::format_date_col`, which renders the value with the user's configured `date_format`.

Two more backend properties are worth setting for custom columns:

* `$date_cols` lists the columns holding a date, so `rcube_addressbook::compare_search_value()` compares them with `rcube_utils::anytodatetime()` instead of doing a substring match.
* `$vcard_map` maps a column to a vCard property, which makes the values survive an export. Use an `X-` extension name, for example `X-VTVBB-MEMBERSHIP-ID`.

### Getting them into the GUI

Declaring the coltype is only half of the job. The contact form is built from the `$form` array in `program/actions/contacts/show.php` (and `edit.php`), and the renderer skips every column which is not listed in a fieldset:

```php
// skip cols not listed in the form definition
if (is_array($fieldset['content']) && !in_array($col, array_keys($fieldset['content']))) {
    continue;
}
```

The `contact_form` hook lets a plugin modify that array before it is rendered, so the plugin adds its own fieldset with the two columns. Note that the `head` section is magic (it only renders the name fields), so custom fields go into a new or an existing regular section.

Labels are a small trap: `run()` translates the labels of `$CONTACT_COLTYPES` *before* the backend definition is merged in, so a `'label' => 'membershipid'` in the backend ends up in the interface verbatim. Resolve the text yourself with the plugin domain, `$rcmail->gettext(['name' => ..., 'domain' => 'static_addressbook'])`, and ship a localization file with the plugin.

## The plugin

Drop the three files below in `plugins/static_addressbook/` and add `static_addressbook` to the `plugins` array in `config/config.inc.php`.

### `plugins/static_addressbook/static_addressbook.php`

```php
<?php

require_once __DIR__ . '/static_addressbook_backend.php';

/**
 * Add a read-only address book with a static list of contacts
 */
class static_addressbook extends rcube_plugin
{
    private $abook_id = 'static';
    private $abook_name = 'Static List';

    public function init()
    {
        // register the plugin texts under the 'static_addressbook' domain
        $this->add_texts('localization/', true);

        $this->add_hook('addressbooks_list', [$this, 'address_sources']);
        $this->add_hook('addressbook_get', [$this, 'get_address_book']);
        $this->add_hook('contact_form', [$this, 'contact_form']);
    }

    /**
     * Add the plugin defined vtvbb_* fields as an extra section of the contact form
     */
    public function contact_form($p)
    {
        // the head section is magic and only holds the name fields, so add our own
        $p['form']['vtvbb'] = [
            'name' => $this->gettext('vtvbb'),
            'content' => [
                'vtvbb_membership_id' => ['size' => 40],
                'vtvbb_member_since' => ['size' => 12],
            ],
        ];

        return $p;
    }

    public function address_sources($p)
    {
        $abook = new static_addressbook_backend($this->abook_name);

        $p['sources'][$this->abook_id] = [
            'id'       => $this->abook_id,
            'name'     => $this->abook_name,
            'readonly' => $abook->readonly,
            'groups'   => $abook->groups,
        ];

        return $p;
    }

    public function get_address_book($p)
    {
        if ($p['id'] === $this->abook_id) {
            $p['instance'] = new static_addressbook_backend($this->abook_name);
        }

        return $p;
    }
}
```

### `plugins/static_addressbook/static_addressbook_backend.php`

```php
<?php

/**
 * Read-only address book backend holding a static list of contacts
 * including a home and a work address per contact and the plugin
 * defined vtvbb_* fields
 */
class static_addressbook_backend extends rcube_addressbook
{
    public $primary_key = 'ID';
    public $readonly = true;
    public $groups = false;

    // Restrict the composite 'address' column to the home and work subtypes.
    // The child fields (street, locality, zipcode, region, country) are
    // inherited from the general Roundcube contact column definition.
    // The vtvbb_* columns are unknown to Roundcube and are therefore
    // defined completely here, labels are filled in by the constructor.
    public $coltypes = [
        'name'         => ['limit' => 1],
        'firstname'    => ['limit' => 1],
        'surname'      => ['limit' => 1],
        'jobtitle'     => ['limit' => 1],
        'organization' => ['limit' => 1],
        'email'        => ['subtypes' => ['home', 'work']],
        'phone'        => ['subtypes' => ['home', 'work']],
        'address'      => ['subtypes' => ['home', 'work'], 'limit' => 2],

        'vtvbb_membership_id' => [
            'type'      => 'text',
            'size'      => 40,
            'maxlength' => 32,
            'limit'     => 1,
            'category'  => 'vtvbb',
        ],
        'vtvbb_member_since' => [
            'type'        => 'date',
            'size'        => 12,
            'maxlength'   => 16,
            'limit'       => 1,
            'category'    => 'vtvbb',
            'render_func' => 'rcmail_action_contacts_index::format_date_col',
        ],
    ];

    // Dates are compared with rcube_utils::anytodatetime() while searching
    public $date_cols = ['vtvbb_member_since'];

    // Export the custom fields as vCard extensions
    public $vcard_map = [
        'vtvbb_membership_id' => 'X-VTVBB-MEMBERSHIP-ID',
        'vtvbb_member_since'  => 'X-VTVBB-MEMBER-SINCE',
    ];

    private $name;
    private $filter;
    private $result;

    private $contacts = [
        [
            'ID'            => '111',
            'name'          => 'John Doe',
            'firstname'     => 'John',
            'surname'       => 'Doe',
            'jobtitle'      => 'Embedded Software Engineer',
            'organization'  => 'Example B.V.',
            'email:work'    => ['john.doe@example.org'],
            'email:home'    => ['john@example.net'],
            'phone:work'    => ['+31 24 1234567'],
            'address:home'  => [
                [
                    'street'   => 'Dorpsstraat 1',
                    'locality' => 'Nijmegen',
                    'zipcode'  => '6511 AA',
                    'region'   => 'Gelderland',
                    'country'  => 'Netherlands',
                ],
            ],
            'address:work'  => [
                [
                    'street'   => 'Industrieweg 42',
                    'locality' => 'Arnhem',
                    'zipcode'  => '6827 BM',
                    'region'   => 'Gelderland',
                    'country'  => 'Netherlands',
                ],
            ],
            'vtvbb_membership_id' => 'VTVBB-000111',
            'vtvbb_member_since'  => '2019-04-02',
        ],
        [
            'ID'            => '112',
            'name'          => 'Jane Example',
            'firstname'     => 'Jane',
            'surname'       => 'Example',
            'organization'  => 'Example B.V.',
            'email:work'    => ['jane.example@example.org'],
            'address:work'  => [
                [
                    'street'   => 'Industrieweg 42',
                    'locality' => 'Arnhem',
                    'zipcode'  => '6827 BM',
                    'country'  => 'Netherlands',
                ],
            ],
            'vtvbb_membership_id' => 'VTVBB-000112',
            'vtvbb_member_since'  => '2023-08-27',
        ],
    ];

    public function __construct($name)
    {
        $this->ready = true;
        $this->name = $name;

        // Roundcube translates the labels of its own columns before the
        // backend definition is merged in, so translate ours here
        $rcmail = rcmail::get_instance();

        foreach (['vtvbb_membership_id', 'vtvbb_member_since'] as $col) {
            $this->coltypes[$col]['label'] = $rcmail->gettext([
                'name' => $col,
                'domain' => 'static_addressbook',
            ]);
        }
    }

    public function get_name()
    {
        return $this->name;
    }

    public function set_search_set($filter)
    {
        $this->filter = $filter;
    }

    public function get_search_set()
    {
        return $this->filter;
    }

    public function reset()
    {
        $this->result = null;
        $this->filter = null;
    }

    public function list_groups($search = null, $mode = 0)
    {
        return [];
    }

    public function list_records($cols = null, $subset = 0, $nocount = false)
    {
        // Note: paging is not implemented, the list is small enough
        return $this->result = $this->count();
    }

    public function count()
    {
        $result = new rcube_result_set(0, ($this->list_page - 1) * $this->page_size);

        foreach ($this->contacts as $contact) {
            $result->add($contact);
            $result->count++;
        }

        return $result;
    }

    public function get_result()
    {
        return $this->result;
    }

    public function search($fields, $value, $mode = 0, $select = true, $nocount = false, $required = [])
    {
        $result = new rcube_result_set();

        if (!is_string($value) || !strlen($value)) {
            return $result;
        }

        foreach ($this->contacts as $contact) {
            if (stripos($this->flatten($contact), $value) !== false) {
                $result->add($contact);
                $result->count++;
            }
        }

        return $this->result = $result;
    }

    public function get_record($id, $assoc = false)
    {
        $result = new rcube_result_set(0);

        foreach ($this->contacts as $contact) {
            if ($contact['ID'] == $id) {
                if ($assoc) {
                    return $contact;
                }

                $result->add($contact);
                $result->count = 1;
            }
        }

        return $result;
    }

    /**
     * Flatten a (nested) contact record into a single searchable string
     */
    private function flatten($data)
    {
        $out = '';

        foreach ((array) $data as $value) {
            $out .= ' ' . (is_array($value) ? $this->flatten($value) : (string) $value);
        }

        return $out;
    }
}
```

### `plugins/static_addressbook/localization/en_US.inc`

```php
<?php

$labels = [];
$labels['vtvbb'] = 'VTVBB';
$labels['vtvbb_membership_id'] = 'Membership ID';
$labels['vtvbb_member_since'] = 'Member since';
```

## Result

Open the addressbook, pick the *Static List* source and the contact details show two blocks under *Address*, labelled `Home` and `Work`. The labels come from Roundcube's localization of the subtype, so they are translated for free. Exporting the contact as vCard yields the expected `ADR;TYPE=home` and `ADR;TYPE=work` properties.

Below the standard sections there is now a *VTVBB* fieldset with *Membership ID* (`VTVBB-000111`) and *Member since*, the latter formatted with the date format from the user settings instead of the raw `2019-04-02` from the record. In the vCard export the two fields show up as `X-VTVBB-MEMBERSHIP-ID` and `X-VTVBB-MEMBER-SINCE`.

A couple of gotchas I ran into:

* Forgetting `address` in `$coltypes` silently hides the addresses; Roundcube intersects its column definition with the backend one.
* The value of `address:home` is a list of address arrays, not a single array. A plain `['street' => ...]` is rendered as one address per child field.
* The child keys must match the ones from the general definition (`street`, `locality`, `zipcode`, `region`, `country`); positional arrays work too, but are unreadable.
* A custom column needs *both* a `$coltypes` entry and a `contact_form` hook; with only one of the two it stays invisible.
* Labels of custom columns are not run through `gettext()` by Roundcube, resolve them yourself.
* `readonly = true` saves you from implementing `insert()`, `update()` and `delete()`. For a writable backend remember that a date field comes back from the form as a localized string, so normalize it with `rcube_utils::anytodatetime()` before storing.
