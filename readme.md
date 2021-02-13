#HN To RSS

HN To RSS is an intentionally simple rss feed of the front page of HackerNews. It returns only story links from the first page - no comments, no Ask HN, nothing else.

## Why?

There are plenty of ways to get HackerNews into your rss reader. There's the [official feed](https://news.ycombinator.com/rss) and [hnrss](https://hnrss.github.io) and there are probably others. Why create a new one? First, I wanted to practice with Go and this was a simple project to get started.

Second, I've been on a self-hosted kick recently. At the same time, I've recognized how much time I spend refreshing sites I read to look for new stories and endlessly scrolling feeds. I noticed on HackerNews and Reddit I frequently go straight to the comments without reading the linked article. Because of this, I want to focus my time on the internet more. I set up [Miniflux](https://miniflux.app) and pulled in some sites I read regularly. To combat going to the comments first, I've made Hn To Rss so it doesn't include comments or text posts at all, just the submitted link.
