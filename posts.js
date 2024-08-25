
// send GET request to fetch feed from backend server
fetch("http://localhost:3333/feed") 
    .then((response) => {
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        return response.json()
    })
    .then((data) => {
        
        data.Channel.Items.forEach((item) => {
            const maxPreviewLength = 300

            const post = document.createElement("div");

            const title = document.createElement("h3");
            title.innerText = item.Title;
            post.appendChild(title);

            const publishDate = document.createElement("p");
            publishDate.innerText = item.PubDate;
            publishDate.style.font = "italic 0.8rem Courier,serif";
            post.appendChild(publishDate);

            const content = document.createElement("div");
            content.innerHTML = (item.Content.length <= maxPreviewLength) ? item.Content : item.Content.substr(0, maxPreviewLength) + '\u2026';
            post.appendChild(content)

            document.getElementById("posts").appendChild(post);
        });
        console.log(data);
    })
    .catch((error) => console.error("Fetch error: ", error));
