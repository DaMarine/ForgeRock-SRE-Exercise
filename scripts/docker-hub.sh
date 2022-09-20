docker build --tag stock-ticker .
docker tag stock-ticker pierswilliams/stock-ticker:1.0.0
docker push pierswilliams/stock-ticker:1.0.0