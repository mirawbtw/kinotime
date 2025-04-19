document.addEventListener("DOMContentLoaded", () => {
    const moviesContainer = document.getElementById("moviesContainer");

    fetch("http://localhost:8080/movies")
        .then(response => response.json())
        .then(data => {
            if (data.movies) {
                data.movies.forEach(movie => {
                    const movieCard = document.createElement("div");
                    movieCard.classList.add("col-md-4", "mb-4");

                    movieCard.innerHTML = `
                        <div class="card shadow-sm">
                            <img src="${movie.poster_url}" class="card-img-top" alt="${movie.title}">
                            <div class="card-body">
                                <h5 class="card-title">${movie.title} (${movie.year})</h5>
                                <p class="card-text"><strong>Genre:</strong> ${movie.genre}</p>
                                <p class="card-text">${movie.description}</p>
                                <p class="card-text"><strong>Actors:</strong> ${movie.actors.join(", ")}</p>
                            </div>
                        </div>
                    `;

                    moviesContainer.appendChild(movieCard);
                });
            }
        })
        .catch(error => console.error("Error fetching movies:", error));
});
