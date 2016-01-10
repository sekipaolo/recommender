# Description

Simple go executable for recommendations based on purchase history.
Datases should have following csv format:

```csv
1331072795;4;TTTTCCCCHHHEHSG2
1331074425;1;SC016FL67BJWSLID
1331306282;12;GG848FL83DOYSLID
1331306282;12;BL152FL82CRXSLID
1331306313;11;CS189FL29SGOSLID
1331306332;11;JD629FL54SNHSLID
1331306341;10;DH756FL65HDYSLID
1331306414;7;UI756FL55HSKSLID
1331306414;7;UI756FL55HSKSLID
1331306414;7;UI756FL55HAKALID

```

# USAGE

```bash
go get github.com/sekipaolo/recommender
go install github.com/sekipaolo/recommender 
$GOPATH/bin/recommender 12 testdata.csv
```