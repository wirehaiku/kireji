# Kireji

**Kireji** is a plaintext note-taking CLI, written in [Go 1.20][gver] by [Stephen Malone][smal].

- See [`changes.md`][chng] for the complete changelog.
- See [`license.md`][lcns] for licensing information (BSD-3).

If you've got a directory full of plaintext notes you need to handle, Kireji can give you an easy-breezy command-line interface to do just that.
Just tell Kireji where your notes are and let it do the rest.

## Installation

You can install Kireji with `go get`...

```
go get github.com/wirehaiku/kireji@latest
```

or download a binary from the [latest release][rels].

## Configuration

Kireji just needs two environment variables to be set and exported.

```fish
# KIREJI_DIR is the path to your notes directory.
KIREJI_DIR = "~/notes"

# KIREJI_EXT is the file extension your notes use.
KIREJI_EXT = ".txt"
```

That's it! That's all the configuration you need.

## Commmands

### Syntax

Kireji represents note files as lowercase base names.
For example, the file `~/notes/foo.txt` is represented as `foo` and asking Kireji to make the note `FOO BAR` results in `~/notes/foo-bar.txt` being created.

### `show NOTE`

Print a note's contents, if it exists:

<summary> <details> Example: </details> 

```
$ kireji show groceries
Need to buy: bread, milk, tomato soup, party streamers...
```

</summary>

## Contributing

Please submit all feature requests and bug reports to the [issue tracker][bugs], thank you.

[bugs]: https://github.com/wirehaiku/kireji/issues
[chng]: https://github.com/wirehaiku/kireji/blob/main/changes.md
[gver]: https://go.dev/doc/go1.20
[lcns]: https://github.com/wirehaiku/kireji/blob/main/license.md
[rels]: https://github.com/wirehaiku/kireji/releases/latest
[smal]: https://wirehaiku.org/


