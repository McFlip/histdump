# histdump

A forensics tool to extract browser history from `sqlite` databases used by Chrome, Edge, and Firefox.

Dumps the history sqlite database into a CSV file for easier analysis and reporting.

The input is the history file and the output is a CSV file.

It's written in Go and is a single executable with no dependencies. It can be run on Windows, Linux, or MacOS.

## Purpose

It's speed and simplicity make it ideal for triage and quick analysis of browser history data. It can also be used in scripts to automate the processing of browser history data.

It's a command line tool but is designed to be user friendly thanks to [Cobra Command](https://cobra.dev/).

Burn it to a DVD along with a copy of [PowerShell](https://github.com/PowerShell/PowerShell) or `Bash` and you have a powerful tool for processing browser history data on the go.

Use the `Import-CSV` cmdlet in PowerShell or `grep` in Bash to quickly search through the CSV file for specific URLs.

## Database Locations

> [!NOTE]
> MS Edge is just Chrome with a MS paint job. The schema for the history files is the same, so you process them in the tool the same.

These are the default locations on Windows. For Linux and MacOS, the locations will be different but the file names will be the same.

### Chrome `History`

```powershell
%LOCALAPPDATA%\Google\Chrome\User Data\Default\
```

### Edge `History`

```powershell
%LOCALAPPDATA%\Microsoft\Edge\User Data\Default
```

### Firefox `places.sqlite`

```powershell
%APPDATA%\Mozilla\Firefox\Profiles\<profile>\places.sqlite
```

`profile` is randomly generated for each profile, so you will have to drill down through the tree.

## Download the tool

You can download the latest release from the [releases page](https://github.com/McFlip/histdump/releases).

It's a single executable with no dependencies, so you can just download it and run it.

If you plan on using it for triage on a live target system, download all the releases so you are prepared for any OS you might encounter.

## Usage

Run once for each browser.

```powershell
.\browser-history.exe [ chrome | firefox ] --file <Input DB> --output <CSV output> --after YYYY-MM-DD --before YYYY-MM-DD
```

chrome
: Use for both Chrome and Edge `History` files

firefox
: Use for Firefox `places.sqlite` file

file
: Path to input (required)

output
: Path to output (required)

after (optional)
: Starting date of investigation. Must be in Year-Month-Day format.

before (optional)
: Ending date of investigation. Time cutoff is 0000 UTC so add 1 day buffer to get events on the last night of your timeframe.

You can specify `after`, `before`, or both together to get a `between` filter.

> [!NOTE]
> You can pass the `--help` flag to the root command, chrome command, or the firefox command.
> You can also just run `.\browser-history.exe` by itself and get usage info
