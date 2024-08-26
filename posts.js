
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

            const title = document.createElement("a");
            title.innerText = item.Title;
            title.style.font = "1.5rem Courier,serif";
            title.setAttribute('href', '#');      
            title.onclick = function(){loadPostDetail(item)};
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

const loadPostDetail = (item) => {
    const posts = document.getElementById("posts");
    posts.innerHTML = "";

    // update dir path at top of page
    const dir = document.getElementById("dir")
    dir.innerHTML = "";

    const backLink = document.createElement("a");
    backLink.innerText = "< Back to list";
    backLink.setAttribute('href', 'posts.html');
    dir.appendChild(backLink);

    const title = document.createElement("h2");
    title.innerText = item.Title;
    posts.appendChild(title);

    const byLine = document.createElement("p");
    byLine.innerText = "Written " + item.PubDate + " by " + item.Creator;
    posts.appendChild(byLine);

    const sourceLink = document.createElement("a");
    sourceLink.innerText = "View Original";
    sourceLink.setAttribute('href', item.Link);
    posts.appendChild(sourceLink);

    const content = document.createElement("div");
    content.innerHTML = item.Content;
    posts.appendChild(content);

    console.log("run");
};
