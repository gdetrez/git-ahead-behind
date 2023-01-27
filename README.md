# git ahead-behind

A small git extension that summarize the relative state of all your brannches.

```sh
$ git ahead-behind
  new-feature           f0dd42d6    0 ┝ 2  
  abandonned            717e4a57  266━┿ 1  
* main                  32b0d398    0 │ 0  
  recent-work           85f50c13    6 ┿╸12 
  no-idea-what-this-was 668a764e  391━┿ 1  
```

## Installation

Install from source using `go install`

```bash
go install github.com/gdetrez/git-ahead-behind@main
```

## Usage

```sh
$ git ahead-behind [(-r | --remotes) | (-a | --all)] [--base <branch>]
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
