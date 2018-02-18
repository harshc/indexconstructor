# indexconstructor

Command line utility to build and serialize a json dataset.

# Usage

This tool has a built in Help function


``` 
    $ indexconstructor --help
    Constructs a JSON dataset for the custom index processor
    
    Usage:
      indexconstructor [flags]
      indexconstructor [command]
    
    Available Commands:
      generate    Generates a json file from a source word list
      help        Help about any command
      process     Process a json file and generate a serialized data structure
    
    Flags:
      -h, --help   help for indexconstructor
    
    Use "indexconstructor [command] --help" for more information about a command.    
```

There are two main actions to this tool
## Generate
This action consumes a txt file with a list of words and generates a formatted json file that can be consumed by the _**Process**_ action.
This particular action is really a helper action to generate a dataset for the processor.
The list of words is in the mock_dataset.yaml file. The reference data set is generated from [SCOWL](http://app.aspell.net/create?max_size=80&max_variant=0&diacritic=strip&special=hacker&download=wordlist&encoding=utf-8&format=inline) and is appropriately licensed.
This action will add a random number score to each word and store it in a json structure.

```
    $ indexconstructor generate --help
    Reads dummy word list from the included yaml file and outputs a json file.
    
    Usage:
      indexconstructor generate [flags]
    
    Flags:
      -h, --help             help for generate
      -f, --input string     name of the input file to read from (default "dataset.txt")
      -o, --outfile string   name of the file to write to (default "dataset.json")

```

## Process
This is the main action of this tool. This processes dataset.json (or an input file) and streams it to an http endpoint

```
    $ indexconstructor process --help
    Process the json file and stream the serialized json to an API endpoint
    
    Usage:
      indexconstructor process [flags]
    
    Flags:
      -h, --help           help for process
      -f, --input string   Input file with the dataset (default "dataset.json")
      -u, --url string     API Url to invoke with the json streaming payload

``` 
