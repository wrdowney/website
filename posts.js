
// send GET request to fetch feed from backend server
fetch("http://localhost:3333/feed") 
    .then((response) => {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }

        return response.json();
    })
    .then((data) => console.log(data))
    .catch((error) => console.error("Fetch error: ", error));
