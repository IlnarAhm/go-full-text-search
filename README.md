# Simple Full-Text Search engine

get the data dump from here - https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz

## How to run

```bash
go run main.go -q "query string" -p "dump path"
```

### Example

run program
```bash
go run main.go -q "In software engineering, a domain model is a conceptual model"
```
result
```bash
Full text search is in progress
Loaded 690680 documents in 19.500479583s
Indexed 690680 documents in 9.503804083s
Search found 1 documents in 12.215875ms
626737      In software engineering, a domain model is a conceptual model of the domain that incorporates both behavior and data.Fowler, Martin.
```