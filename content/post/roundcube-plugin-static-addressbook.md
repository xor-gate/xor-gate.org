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

Roundcube ships with an [`example_addressbook`](https://github.com/roundcube/roundcubemail/tree/master/plugins/example_addressbook) plugin which exposes a hardcoded list of contacts as an extra address source. It is a great starting point when you want to expose an internal directory (LDAP is not always the answer) inside the webmail. What the upstream example does not show is how to return *structured* fields, like a postal address with a `home` and `work` variant.

Below is a small, complete plugin which does exactly that: a read-only addressbook with static contacts, each carrying a home and a work address.

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

## The plugin

Drop the two files below in `plugins/static_addressbook/` and add `static_addressbook` to the `plugins` array in `config/config.inc.php`.

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
        $this->add_hook('addressbooks_list', [$this, 'address_sources']);
        $this->add_hook('addressbook_get', [$this, 'get_address_book']);
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
 * including a home and a work address per contact
 */
class static_addressbook_backend extends rcube_addressbook
{
    public $primary_key = 'ID';
    public $readonly = true;
    public $groups = false;

    // Restrict the composite 'address' column to the home and work subtypes.
    // The child fields (street, locality, zipcode, region, country) are
    // inherited from the general Roundcube contact column definition.
    public $coltypes = [
        'name'         => ['limit' => 1],
        'firstname'    => ['limit' => 1],
        'surname'      => ['limit' => 1],
        'jobtitle'     => ['limit' => 1],
        'organization' => ['limit' => 1],
        'email'        => ['subtypes' => ['home', 'work']],
        'phone'        => ['subtypes' => ['home', 'work']],
        'address'      => ['subtypes' => ['home', 'work'], 'limit' => 2],
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
        ],
    ];

    public function __construct($name)
    {
        $this->ready = true;
        $this->name = $name;
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

## Result

Open the addressbook, pick the *Static List* source and the contact details show two blocks under *Address*, labelled `Home` and `Work`. The labels come from Roundcube's localization of the subtype, so they are translated for free. Exporting the contact as vCard yields the expected `ADR;TYPE=home` and `ADR;TYPE=work` properties.

A couple of gotchas I ran into:

* Forgetting `address` in `$coltypes` silently hides the addresses; Roundcube intersects its column definition with the backend one.
* The value of `address:home` is a list of address arrays, not a single array. A plain `['street' => ...]` is rendered as one address per child field.
* The child keys must match the ones from the general definition (`street`, `locality`, `zipcode`, `region`, `country`); positional arrays work too, but are unreadable.
* `readonly = true` saves you from implementing `insert()`, `update()` and `delete()`.
