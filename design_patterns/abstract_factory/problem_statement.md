You are building a cross-platform UI toolkit.

The application must support:

Windows Theme
```
Windows Button
Windows Checkbox
```

Mac Theme
```
Mac Button
Mac Checkbox
```

Every button supports:
```Go
Render()
```

Every checkbox supports:
```Go
Render()
```
Current Requirement

When the application starts:
```
theme = "windows"
```
or
```
theme = "mac"
```

and the entire UI should be created using components from that theme.

Examples:
```
Windows Button + Windows Checkbox
```
or
```
Mac Button + Mac Checkbox
```

Mixing is not allowed.


This is invalid:
```
Windows Button + Mac Checkbox
```

## Future Requirement

Next quarter:

Linux Theme

will be added.

That means:
```
Linux Button
Linux Checkbox
```