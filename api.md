# Shared Response Types

- <a href="https://pkg.go.dev/github.com/Influship/influship-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go/shared#CreatorBasic">CreatorBasic</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go/shared#ProfileSummary">ProfileSummary</a>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/Influship/influship-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Creators

Response Types:

- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorGetResponse">CreatorGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorAutocompleteResponse">CreatorAutocompleteResponse</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorLookalikeResponse">CreatorLookalikeResponse</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorMatchResponse">CreatorMatchResponse</a>

Methods:

- <code title="get /v1/creators/{id}">client.Creators.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorGetParams">CreatorGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorGetResponse">CreatorGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/creators/autocomplete">client.Creators.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorService.Autocomplete">Autocomplete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorAutocompleteParams">CreatorAutocompleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorAutocompleteResponse">CreatorAutocompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/creators/lookalike">client.Creators.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorService.Lookalike">Lookalike</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorLookalikeParams">CreatorLookalikeParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go/packages/pagination#BodyCursor">BodyCursor</a>[<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorLookalikeResponse">CreatorLookalikeResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/creators/match">client.Creators.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorService.Match">Match</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorMatchParams">CreatorMatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#CreatorMatchResponse">CreatorMatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Search

Response Types:

- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#MatchInfo">MatchInfo</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#SearchNewResponse">SearchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#SearchGetResponse">SearchGetResponse</a>

Methods:

- <code title="post /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/Influship/influship-go#SearchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#SearchNewParams">SearchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#SearchNewResponse">SearchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/search/{id}">client.Search.<a href="https://pkg.go.dev/github.com/Influship/influship-go#SearchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#SearchGetParams">SearchGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go/packages/pagination#QueryCursor">QueryCursor</a>[<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#SearchGetResponse">SearchGetResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileActivity">ProfileActivity</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileGrowth">ProfileGrowth</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileMetrics">ProfileMetrics</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileResponseData">ProfileResponseData</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileGetResponse">ProfileGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileLookupResponse">ProfileLookupResponse</a>

Methods:

- <code title="get /v1/profiles/{platform}/{username}">client.Profiles.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileGetParams">ProfileGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileGetResponse">ProfileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/profiles/lookup">client.Profiles.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileLookupParams">ProfileLookupParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#ProfileLookupResponse">ProfileLookupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Posts

Response Types:

- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#PostListResponse">PostListResponse</a>

Methods:

- <code title="get /v1/posts">client.Posts.<a href="https://pkg.go.dev/github.com/Influship/influship-go#PostService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#PostListParams">PostListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go/packages/pagination#QueryCursor">QueryCursor</a>[<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#PostListResponse">PostListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Raw

## Instagram

Response Types:

- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawInstagramGetProfileResponse">RawInstagramGetProfileResponse</a>

Methods:

- <code title="get /v1/raw/instagram/profile/{username}">client.Raw.Instagram.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawInstagramService.GetProfile">GetProfile</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawInstagramGetProfileParams">RawInstagramGetProfileParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawInstagramGetProfileResponse">RawInstagramGetProfileResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Youtube

Response Types:

- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#TranscriptSegment">TranscriptSegment</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeGetChannelResponse">RawYoutubeGetChannelResponse</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeGetChannelTranscriptsResponse">RawYoutubeGetChannelTranscriptsResponse</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeGetTranscriptResponse">RawYoutubeGetTranscriptResponse</a>
- <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeSearchResponse">RawYoutubeSearchResponse</a>

Methods:

- <code title="get /v1/raw/youtube/channel/{handle}">client.Raw.Youtube.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeService.GetChannel">GetChannel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, handle <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeGetChannelParams">RawYoutubeGetChannelParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeGetChannelResponse">RawYoutubeGetChannelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/raw/youtube/channel-transcripts/{handle}">client.Raw.Youtube.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeService.GetChannelTranscripts">GetChannelTranscripts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, handle <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeGetChannelTranscriptsParams">RawYoutubeGetChannelTranscriptsParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeGetChannelTranscriptsResponse">RawYoutubeGetChannelTranscriptsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/raw/youtube/transcript/{video_id}">client.Raw.Youtube.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeService.GetTranscript">GetTranscript</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, videoID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeGetTranscriptParams">RawYoutubeGetTranscriptParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeGetTranscriptResponse">RawYoutubeGetTranscriptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/raw/youtube/search">client.Raw.Youtube.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeSearchParams">RawYoutubeSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/Influship/influship-go">influshipapi</a>.<a href="https://pkg.go.dev/github.com/Influship/influship-go#RawYoutubeSearchResponse">RawYoutubeSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
