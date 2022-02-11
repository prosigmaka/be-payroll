# Build image..
docker build -t aditya/bePayroll --no-cache .

# 2.
# run and create container
# with open port
docker run -d --name=bepayroll -p 8081:8081 aditya/bePayroll

# diferent port
#docker run -d --name=helloapp2 -p 5000:3000 arnugroho/hello