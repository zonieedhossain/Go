package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"reflect"

	"github.com/twpayne/go-polyline"
)

type pointlist struct{ x, y float64 }

func perpendicularDistance(p, p1, p2 pointlist) (r float64) {
	if p1.x == p2.x {
		r = math.Abs(float64(p.x - p1.x))
	} else {
		slope := float64(p2.y-p1.y) / float64(p2.x-p1.x)
		intercept := float64(p1.y) - (slope * float64(p1.x))
		r = math.Abs(slope*float64(p.x)-float64(p.y)+intercept) / math.Sqrt(math.Pow(slope, 2)+1)
	}
	return
}

func RDP(list []pointlist, epsilon float64) []pointlist {
	if len(list) < 3 {
		return list
	}
	dmax := 0.0
	index := 0
	end := len(list) - 1
	for i := 2; i < len(list)-1; i++ {
		d := perpendicularDistance(list[i], list[1], list[end])
		if d > dmax {
			index = i
			dmax = d
		}
	}
	if dmax > epsilon {
		return append(RDP(list[:index+1], epsilon), RDP(list[index:], epsilon)[1:]...)
	}
	return []pointlist{list[0], list[len(list)-1]}

}
func GetMapWithPolygon(output string, geometry string) {
	endpoint, _ := url.Parse("https://maps.googleapis.com/maps/api/staticmap?center=Dhaka&size=2480x720&path=weight:2%7Ccolor:red%7Cenc:}swpC_jsfPo@jSnIpK|OsO&path=weight:2%7Ccolor:blue%7Cenc:_`lpCmlxfPnWwdAnY_hAxHqqA&key=YOUR_API_KEY")
	queryParams := endpoint.Query()
	queryParams.Set("ppi", "320")
	queryParams.Set("w", "1280")
	queryParams.Set("h", "720")
	queryParams.Set("z", "1")
	queryParams.Set("a0", geometry)
	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())
	if err != nil {
		fmt.Printf("the http request faild %s\n", err)
	} else {
		f, err := os.Create(output)
		if err != nil {
			fmt.Printf("the http request faild %s\n", err)

			return
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("the http request faild %s\n", err)
			return
		}
		f.Write(data)
		defer f.Close()
	}
}
func main() {
	r := RDP([]pointlist{{23.879519, 90.396992}, {23.878037, 90.397432}, {23.878067, 90.398263}, {23.878121, 90.398965}, {23.877684, 90.399035}, {23.876733, 90.399265}, {23.875933, 90.399710}}, 0)
	fmt.Println(reflect.TypeOf(r))
	e := polyline.EncodeCoord(r)
	fmt.Println(e)

	GetMapWithPolygon("map.jpeg", "r")
}
