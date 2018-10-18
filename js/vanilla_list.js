const createVanillaChildEl = index => {
    const child = document.createElement("li");
    const span = document.createElement("span");
    span.innerText = "Child Element " + index;
    child.appendChild(span);
    return child
}

const renderVanilla = count => {
    const listContainer = document.getElementById('vanilla-list-demo-container');
    const flameChart = document.createElement("img");
    const vanillaListContainer = document.getElementById('vanilla-flame-chart-container');

    flameChart.setAttribute("src", "./img/vanilla-flame-chart.png");
    vanillaListContainer.appendChild(flameChart)

    for (let i = 1; i <= count; i++) {
        listContainer.appendChild(createVanillaChildEl(i))
    }
}