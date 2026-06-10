// DOM Elements
const cityInput = document.getElementById('city-input');
const searchBtn = document.getElementById('search-btn');
const favoriteBtn = document.getElementById('favorite-btn');
const favoritesList = document.getElementById('favorites-list');
const weatherResult = document.getElementById('weather-result');
const errorMsg = document.getElementById('error-msg');
const loader = document.getElementById('loader');

const cityName = document.getElementById('city-name');
const temp = document.getElementById('temp');
const description = document.getElementById('description');
const feelsLike = document.getElementById('feels-like');
const humidity = document.getElementById('humidity');
const wind = document.getElementById('wind');
const rainProb = document.getElementById('rain-prob');
const sunrise = document.getElementById('sunrise');
const sunset = document.getElementById('sunset');

let favorites = JSON.parse(localStorage.getItem('weather-favorites') || '[]');

// Weather specific effects
function applyWeatherEffects(condition, isNight, isGoldenHour) {
    document.body.className = '';
    document.body.classList.add(condition);
    if (isNight) document.body.classList.add('night');
    if (isGoldenHour) document.body.classList.add('golden-hour');

    document.querySelectorAll('.rain-drop').forEach(el => el.remove());

    if (condition === 'Rain' || condition === 'Drizzle') {
        for (let i = 0; i < 30; i++) {
            const drop = document.createElement('div');
            drop.className = 'rain-drop';
            drop.style.left = Math.random() * 100 + 'vw';
            drop.style.animationDuration = (Math.random() * 0.5 + 0.5) + 's';
            drop.style.opacity = Math.random();
            document.body.appendChild(drop);
        }
    }
}

function updateForecastChart(forecastData) {
    const chartContainer = document.getElementById('chart-container');
    if (!chartContainer || !forecastData) return;
    chartContainer.innerHTML = '';
    
    const points = forecastData.slice(0, 8);
    const temperatures = points.map(p => p.Temp);
    const maxTemp = Math.max(...temperatures);
    const minTemp = Math.min(...temperatures);
    const range = maxTemp - minTemp || 1;

    points.forEach(point => {
        const bar = document.createElement('div');
        bar.className = 'chart-bar';
        const height = ((point.Temp - minTemp) / range) * 70 + 30;
        bar.style.height = `${height}%`;
        bar.innerHTML = `
            <span class="bar-temp">${Math.round(point.Temp)}°</span>
            <span class="bar-time">${point.Time}</span>
        `;
        chartContainer.appendChild(bar);
    });
}

function renderFavorites() {
    favoritesList.innerHTML = '';
    favorites.forEach(city => {
        const pill = document.createElement('div');
        pill.className = 'fav-pill';
        pill.textContent = city;
        pill.onclick = () => {
            cityInput.value = city;
            getWeather(city);
        };
        // Right click to remove
        pill.oncontextmenu = (e) => {
            e.preventDefault();
            favorites = favorites.filter(f => f !== city);
            localStorage.setItem('weather-favorites', JSON.stringify(favorites));
            renderFavorites();
        };
        favoritesList.appendChild(pill);
    });
}

async function getWeather(specificCity = null) {
    const city = specificCity || cityInput.value.trim();
    if (!city) return;

    weatherResult.classList.add('hidden');
    errorMsg.classList.add('hidden');
    loader.classList.remove('hidden');

    try {
        const result = await window.go.main.App.GetWeather(city);
        
        cityName.textContent = result.City;
        temp.textContent = Math.round(result.Temperature);
        description.textContent = result.Description;
        feelsLike.textContent = `${result.FeelsLike.toFixed(1)}°C`;
        humidity.textContent = `${result.Humidity}%`;
        wind.textContent = `${result.WindSpeed} m/s`;
        rainProb.textContent = `${result.RainProb}%`;
        sunrise.textContent = result.Sunrise;
        sunset.textContent = result.Sunset;

        applyWeatherEffects(result.Condition, result.IsNight, result.IsGoldenHour);
        updateForecastChart(result.Forecast);
        weatherResult.classList.remove('hidden');
    } catch (err) {
        errorMsg.textContent = err;
        errorMsg.classList.remove('hidden');
        document.body.className = 'default';
    } finally {
        loader.classList.add('hidden');
    }
}

async function init() {
    renderFavorites();
    try {
        const localCity = await window.go.main.App.GetLocalCity();
        if (localCity) {
            cityInput.value = localCity;
            getWeather(localCity);
        }
    } catch (e) {
        console.log("Could not detect local city");
    }
}

favoriteBtn.onclick = () => {
    const city = cityName.textContent;
    if (city && city !== '--' && !favorites.includes(city)) {
        favorites.push(city);
        localStorage.setItem('weather-favorites', JSON.stringify(favorites));
        renderFavorites();
    }
};

searchBtn.onclick = () => getWeather();
cityInput.onkeypress = (e) => { if (e.key === 'Enter') getWeather(); };

// Initialize
init();
