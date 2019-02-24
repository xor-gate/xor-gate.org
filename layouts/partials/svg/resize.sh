#!/bin/sh
size=32
for file in *.svg;
do
	echo "resize $file"
	tmpfile=$file.$size
	rsvg-convert -a -w $size -f svg $file -o $tmpfile
	# chomp of the xml header
	sed '1d' $tmpfile > $file
	rm $tmpfile
done
