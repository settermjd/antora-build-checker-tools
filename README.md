# antora-build-checker-tools

This is a small Go app, designed to help with the Antora build process at ownCloud.

At this stage, the ambitions of the app have not been that well defined, however, a start has been made.
Currently, it only checks an AsciiDoc file for a limited amount of reStructuredText formatting.
Specifically, this is only headers.
Formatting such as source includes, lists etc, is not yet covered.

## Installing

To install the app, compile it for your platform.

## Usage

The script, currently, reads a filename from stdin.

### Example

```
echo ./docs/modules/ROOT/pages/configuration/server/antivirus_configuration.adoc | antora-build-checker-tools
```
