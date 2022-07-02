const API_URL = 'http://localhost:3000/api/';

async function get(url) {
    const options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": "*",
        }
    }

    return fetch(url, options).then(res => res.json());
}

async function getInfo() {
    const url = `${API_URL}info`;
    const data = await get(url);
    return data;
}

export default { getInfo, get };