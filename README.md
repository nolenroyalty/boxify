# boxify

Put a box around text.

## Install

```
go install github.com/nolenroyalty/boxify@latest
```

## Usage

```
echo "hello world" | boxify
boxify somefile.txt
```

```
╭─────────────╮
│ hello world │
╰─────────────╯
```

### Options

`--border STYLE` — border style (default: `rounded`). Available styles: `normal`, `rounded`, `block`, `outer-half`, `inner-half`, `thick`, `double`, `hidden`, `markdown`, `ascii`.

`--padding N` or `--padding Y,X` — padding inside the border (default: `0,1`). A single value sets both axes; two comma-separated values set vertical and horizontal separately.

### Examples

```
echo "fancy" | boxify --border double
╔═══════╗
║ fancy ║
╚═══════╝

echo "roomy" | boxify --padding 2,4
╭─────────────╮
│             │
│             │
│    roomy    │
│             │
│             │
╰─────────────╯
```
