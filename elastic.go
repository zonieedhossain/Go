package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	elastic "gopkg.in/olivere/elastic.v5"
)

type Group struct {
	User    string                `json:"user,omitempty"`
	Massege string                `json:"massege,omitempty"`
	Posts   int                   `json:"posts,omitempty"`
	Suggest *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 2,
		"number_of_replicas": 0
	},
	"mappings":{
		"group":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"massege":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

func main() {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	esver, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ES version %s\n", esver)

	exists, err := client.IndexExists("facebook").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("exists\n", exists)
	if !exists {
		createIndex, err := client.CreateIndex("facebook").BodyJson(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println("create\n", createIndex)
		if !createIndex.Acknowledged {
			fmt.Println("not acknowledged")
		}
	}
	group1 := Group{User: "zonieed", Massege: "Thumbs Up", Posts: 2}
	put1, err := client.Index().Index("facebook").Type("group").Id("12").BodyJson(group1).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Index group %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	group2 := `{"user" : "zonieed", "message" : "It's a Raggy Waltz"}`
	put2, err := client.Index().
		Index("facebook").
		Type("group").
		Id("2").
		BodyString(group2).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Index group %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
	get1, err := client.Get().Index("facebook").Type("group").Id("12").Do(ctx)
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("File Got %s version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
	_, err = client.Flush().Index("facebook").Do(ctx)
	if err != nil {
		panic(err)
	}
	if c := os.Getenv("INDEX"); c == "1" {
		for {
			d, err := client.Get().Index("facebook").Type("group").Id("12").Do(ctx)
			if err != nil {
				panic(err)
			}
			fmt.Printf("---------------", d)
		}
	}
	if c := os.Getenv("SEARCH"); c == "1" {
		for {
			t := elastic.NewTermQuery("user", "zonieed")
			r, err := client.Search().Index("facebook").Query(t).Sort("user", true).From(0).Size(10).Pretty(true).Do(ctx)
			if err != nil {
				panic(err)
			}
			fmt.Println("....search....", r)
		}
	}
	termQuery := elastic.NewTermQuery("user", "zonieed")
	fmt.Println(termQuery)
	searchResult, err := client.Search().Index("facebook").Query(termQuery).Sort("user", true).From(0).Size(10).Pretty(true).Do(ctx)
	fmt.Println(searchResult)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	var grp Group
	for _, item := range searchResult.Each(reflect.TypeOf(grp)) {
		if t, ok := item.(Group); ok {
			fmt.Printf("Group by %s: %s\n", t.User, t.Massege)
		}
	}
	fmt.Printf("Found a total of %d groups\n", searchResult.TotalHits())
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d groups\n", searchResult.Hits.TotalHits)

		for _, hit := range searchResult.Hits.Hits {
			var t Group
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				fmt.Println("Deserialization failed")
			}

			fmt.Printf("Group by %s: %s\n", t.User, t.Massege)
		}
	} else {
		fmt.Print("Found no Groups\n")
	}
	if g := os.Getenv("UPDATE"); g == "1" {
		for {
			script := elastic.NewScript("ctx._source.posts += params.num").Param("num", 1)
			update, err := client.Update().Index("facebook").Type("group").Id("1").
				Script(script).
				Upsert(map[string]interface{}{"posts": 0}).
				Do(context.Background())
			if err != nil {
				panic(err)
			}
			fmt.Printf("New version of group %q is now %d\n", update.Id, update.Version)
		}
	}

	deleteIndex, err := client.DeleteIndex("facebook").Do(ctx)
	if err != nil {

		panic(err)
	}
	if !deleteIndex.Acknowledged {
	}
}
