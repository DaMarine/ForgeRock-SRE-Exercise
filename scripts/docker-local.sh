docker build --tag stock-ticker .
docker tag stock-ticker pierswilliams/stock-ticker:1.0.0
docker run -e APIKEY=C227WD9W3LUVKVV9 -e SYMBOL=MSFT -e NDAYS=7 stock-ticker