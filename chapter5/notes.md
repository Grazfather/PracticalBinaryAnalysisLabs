# Chapter 5 notes

* A few binary tools have demangling options that I was unaware of (I usually just pipe to `c++filt`).
  * `nm --demangle`
  * `ltrace -C`
* `ldd` has a `-z` option that will decompress gzip, so you can 'peek into' a gzipped file.
  * Though if it's a `tar.gz`, you'll only see the TAR file, not a layer deeper (e.g. the headers of all the files).
* `xxd` has a `-s` option that will seek to a specified offset.
* To get the length of an ELF file by it's ELF header, look at the 'start of
  section headers' offset, and add to that the 'number of section headers'
  times the 'size of section headers'. This is the end offset.
* `nm` has a `-D` option to look at the _dynamic_ symbols instead of the static ones.
* `strings` has a `-d` option to only look in the data sections of a binary for strings.
