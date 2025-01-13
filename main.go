package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/IlnarAhm/go-full-text-search/utils"
)

func main() {
	var dumpPath, query string

	flag.StringVar(&dumpPath, "p", "enwiki.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	log.Println("Full text search is in progress")
	
	utils.RunPipeline()
	defer utils.ClosePipeline()

	start := time.Now()

	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()

	idx := make(utils.Index)
	idx.Add(docs)

	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()

	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
