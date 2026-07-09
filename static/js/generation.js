document.addEventListener("DOMContentLoaded", function() {
    const searchInputs = document.querySelectorAll(".pokemon-search-input");
    const grid = document.getElementById("pokemon-grid");
    const cards = Array.from(grid.querySelectorAll(".pokemon-card-wrapper"));
    const noResults = document.getElementById("no-results");
    const clearBtn = document.getElementById("clear-search");
    const showingCount = document.getElementById("showing-count");
    const resetFiltersBtn = document.getElementById("reset-filters-btn");
    const typeFilterContainer = document.getElementById("type-filter-container");
    const activeFiltersList = document.getElementById("active-filters-list");

    let currentSearch = "";
    let currentTypes = []; // Array of active type shortnames

    // 1. Gather all unique types in the DOM and map them to their localized name
    const typeMap = {};
    cards.forEach(card => {
        const badges = card.querySelectorAll(".type-badge");
        badges.forEach(badge => {
            const classes = Array.from(badge.classList);
            const typeClass = classes.find(c => c.startsWith("type-") && c !== "type-badge");
            if (typeClass) {
                const shortname = typeClass.replace("type-", "");
                const name = badge.textContent.trim();
                typeMap[shortname] = name;
            }
        });
    });

    // 2. Generate type buttons dynamically
    Object.keys(typeMap).sort().forEach(shortname => {
        const btn = document.createElement("button");
        btn.className = "btn btn-sm btn-outline-secondary text-light m-1 type-filter-btn";
        btn.style.borderRadius = "6px";
        btn.style.borderColor = "rgba(255,255,255,0.15)";
        btn.setAttribute("data-type", shortname);

        // Add a small styled color dot or text representation
        btn.innerHTML = `<span class="d-inline-block mr-1" style="width: 8px; height: 8px; border-radius: 50%; background-color: var(--type-color-${shortname});"></span>${typeMap[shortname]}`;

        typeFilterContainer.appendChild(btn);
    });

    const typeButtons = typeFilterContainer.querySelectorAll(".type-filter-btn");

    // 3. Filter handler function
    function applyFilters() {
        let visibleCount = 0;
        const normalizedSearch = currentSearch.toLowerCase().normalize("NFD").replace(/[\u0300-\u036f]/g, "").trim();

        cards.forEach(card => {
            const name = card.getAttribute("data-name").toLowerCase().normalize("NFD").replace(/[\u0300-\u036f]/g, "");
            const id = card.getAttribute("data-id");
            const types = card.getAttribute("data-types").split(" ").filter(t => t);

            const matchesSearch = !normalizedSearch || name.includes(normalizedSearch) || id.includes(normalizedSearch);
            const matchesType = currentTypes.length === 0 || currentTypes.some(type => types.includes(type));

            if (matchesSearch && matchesType) {
                card.classList.remove("d-none");
                visibleCount++;
            } else {
                card.classList.add("d-none");
            }
        });

        // Update counts
        showingCount.textContent = visibleCount;

        // Show/hide empty state
        if (visibleCount === 0) {
            noResults.classList.remove("d-none");
            grid.classList.add("d-none");
        } else {
            noResults.classList.add("d-none");
            grid.classList.remove("d-none");
        }

        // Update active filters list UI
        activeFiltersList.innerHTML = "";

        let hasFilters = false;

        // 1. Add search query tag
        if (currentSearch.trim()) {
            hasFilters = true;
            const searchTag = document.createElement("span");
            searchTag.className = "badge badge-secondary px-2 py-1 d-inline-flex align-items-center";
            searchTag.style.fontSize = "0.75rem";
            searchTag.style.borderRadius = "6px";
            searchTag.style.backgroundColor = "rgba(255, 255, 255, 0.08)";
            searchTag.style.border = "1px solid rgba(255, 255, 255, 0.1)";
            searchTag.innerHTML = `
                <span class="mr-1 opacity-70">🔍</span> "${currentSearch.trim()}"
                <span class="ml-2 text-warning font-weight-bold" style="cursor: pointer; font-size: 1rem; line-height: 0.8;" id="remove-search-filter">&times;</span>
            `;
            activeFiltersList.appendChild(searchTag);

            searchTag.querySelector("#remove-search-filter").addEventListener("click", () => {
                currentSearch = "";
                searchInputs.forEach(input => input.value = "");
                applyFilters();
            });
        }

        // 2. Add active type tags
        currentTypes.forEach(type => {
            hasFilters = true;
            const typeTag = document.createElement("span");
            typeTag.className = "badge px-2 py-1 d-inline-flex align-items-center";
            typeTag.style.fontSize = "0.75rem";
            typeTag.style.borderRadius = "6px";
            typeTag.style.color = `var(--type-color-${type})`;
            typeTag.style.backgroundColor = `var(--type-bg-${type})`;
            typeTag.style.borderColor = `var(--type-border-${type})`;
            typeTag.style.borderStyle = "solid";
            typeTag.style.borderWidth = "1px";
            typeTag.innerHTML = `
                <span class="d-inline-block mr-1" style="width: 6px; height: 6px; border-radius: 50%; background-color: var(--type-color-${type});"></span>
                ${typeMap[type]}
                <span class="ml-2 font-weight-bold" style="cursor: pointer; font-size: 1rem; line-height: 0.8; opacity: 0.8;" data-remove-type="${type}">&times;</span>
            `;
            activeFiltersList.appendChild(typeTag);

            typeTag.querySelector(`[data-remove-type="${type}"]`).addEventListener("click", () => {
                const index = currentTypes.indexOf(type);
                if (index > -1) {
                    currentTypes.splice(index, 1);
                    // Update type button states in container
                    typeButtons.forEach(b => {
                        if (b.getAttribute("data-type") === type) {
                            b.classList.remove("btn-light");
                            b.classList.add("btn-outline-secondary", "text-light");
                            b.style.borderColor = "rgba(255,255,255,0.15)";
                        }
                    });

                    // Update "All" button
                    const allBtn = Array.from(typeButtons).find(b => b.getAttribute("data-type") === "all");
                    if (currentTypes.length === 0 && allBtn) {
                        allBtn.classList.remove("btn-outline-secondary", "text-light");
                        allBtn.classList.add("btn-light");
                        allBtn.style.borderColor = "transparent";
                    }
                    applyFilters();
                }
            });
        });

        // 3. Fallback if no filters
        if (!hasFilters) {
            const noFilterTag = document.createElement("span");
            noFilterTag.className = "text-muted font-italic";
            noFilterTag.textContent = activeFiltersList.getAttribute("data-no-filters-text") || "";
            activeFiltersList.appendChild(noFilterTag);
        }
    }

    // 4. Bind Search Input Events (synchronizing navbar search & page search)
    searchInputs.forEach(input => {
        input.addEventListener("input", function(e) {
            currentSearch = e.target.value;
            // Sync all inputs
            searchInputs.forEach(otherInput => {
                if (otherInput !== e.target) {
                    otherInput.value = currentSearch;
                }
            });
            applyFilters();
        });
    });

    // 5. Bind clear search button
    if (clearBtn) {
        clearBtn.addEventListener("click", function() {
            currentSearch = "";
            searchInputs.forEach(input => input.value = "");
            applyFilters();
        });
    }

    // 6. Bind Type Filter Buttons
    typeFilterContainer.addEventListener("click", function(e) {
        const btn = e.target.closest(".type-filter-btn");
        if (!btn) return;

        const clickedType = btn.getAttribute("data-type");

        if (clickedType === "all") {
            // Clicked "Tous" / "All"
            currentTypes = [];
            typeButtons.forEach(b => {
                if (b.getAttribute("data-type") === "all") {
                    b.classList.remove("btn-outline-secondary", "text-light");
                    b.classList.add("btn-light");
                    b.style.borderColor = "transparent";
                } else {
                    b.classList.remove("btn-light");
                    b.classList.add("btn-outline-secondary", "text-light");
                    b.style.borderColor = "rgba(255,255,255,0.15)";
                }
            });
        } else {
            // Clicked a specific type
            const index = currentTypes.indexOf(clickedType);
            if (index > -1) {
                // Toggle off
                currentTypes.splice(index, 1);
                btn.classList.remove("btn-light");
                btn.classList.add("btn-outline-secondary", "text-light");
                btn.style.borderColor = "rgba(255,255,255,0.15)";
            } else {
                // Toggle on
                currentTypes.push(clickedType);
                btn.classList.remove("btn-outline-secondary", "text-light");
                btn.classList.add("btn-light");
                btn.style.borderColor = "transparent";
            }

            // Update "All" button active state based on whether any other buttons are selected
            const allBtn = Array.from(typeButtons).find(b => b.getAttribute("data-type") === "all");
            if (currentTypes.length === 0) {
                if (allBtn) {
                    allBtn.classList.remove("btn-outline-secondary", "text-light");
                    allBtn.classList.add("btn-light");
                    allBtn.style.borderColor = "transparent";
                }
            } else {
                if (allBtn) {
                    allBtn.classList.remove("btn-light");
                    allBtn.classList.add("btn-outline-secondary", "text-light");
                    allBtn.style.borderColor = "rgba(255,255,255,0.15)";
                }
            }
        }

        applyFilters();
    });

    // 7. Bind Reset Filters
    function resetAll() {
        currentSearch = "";
        currentTypes = [];
        searchInputs.forEach(input => input.value = "");

        // Reset type buttons state
        typeButtons.forEach(b => {
            if (b.getAttribute("data-type") === "all") {
                b.classList.remove("btn-outline-secondary", "text-light");
                b.classList.add("btn-light");
                b.style.borderColor = "transparent";
            } else {
                b.classList.remove("btn-light");
                b.classList.add("btn-outline-secondary", "text-light");
                b.style.borderColor = "rgba(255,255,255,0.15)";
            }
        });

        applyFilters();
    }

    if (resetFiltersBtn) {
        resetFiltersBtn.addEventListener("click", resetAll);
    }
    
    // Initial application of filters
    applyFilters();
});
