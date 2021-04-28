# webcrawler1

Write a simple web crawler. Given a starting URL, the crawler should visit each URL
it finds on the same domain. It should print each URL visited, and a list of links
found on that page. The crawler should be limited to one subdomain - so when you start
with `https://example.com/`, it would crawl all pages on the `example.com` website,
but not follow external links, for example to `facebook.com` or `community.example.com`.

## Running

```
go build -o crawler .
./crawler -url https://example.com/
```
