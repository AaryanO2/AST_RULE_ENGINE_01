Directions to run
1. Open directory in terminal
2. RUN commands
    docker build -t my-go-app .
    docker run -d -p 8000:8000 --name go-server my-go-app
3. Open browser visit http://localhost:8000/
