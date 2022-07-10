import { Auth } from ".";
import Dummy from '../dummy/data.json';
// const API_URL = (key) => `https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=${key}`;
const HOST_URL = 'http://localhost:3000/';

async function base(url, override={}) {
    let token = Auth.getToken();

    const options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": "*",
            "Authorization": `Bearer ${token ?? ''}`
        },
        ...override
    }

    return fetch(url, options).then(res => res.json()).catch(err => {
        console.log(err, url);
        return false;
    });
}

async function login(Name, Password) {
    const url = HOST_URL + 'api/login';

    return base(url, {
        method: 'POST',
        body: JSON.stringify({
            Name,
            Password
        })
    }).then(data => {
        if (data?.Token) {
            Auth.login(data.Token);
        }
        return data;
    })
}

async function signup(Name, Password) {
    const url = HOST_URL + 'api/signup';

    return base(url, {
        method: 'POST',
        body: JSON.stringify({
            Name,
            Password
        })
    }).then(data => {
        if (data?.Token) {
            Auth.login(data.Token);
        }
        return data;
    })
}

async function getAPIData() {
    const key = Auth.getProfile().APIKey;
    if (!key) return Promise.reject('No API key');
    return Dummy;
    // const url = API_URL(key);
    // return base(url).then(data => data?.result_list);
}

async function setAPIKey(key) {
    const url = HOST_URL + 'api/setAPIKey';

    return base(url, {
        method: 'POST',
        body: JSON.stringify({
            key
        })
    }).then(data => {
        if (data?.Token) {
            Auth.login(data.Token);
        }
        return data;
    })
}

async function getPreferences() {
    const url = HOST_URL + 'api/preferences';

    return base(url);
}

async function setPreferences(prefs) {
    const url = HOST_URL + 'api/preferences';

    return base(url, {
        method: 'POST',
        body: JSON.stringify(prefs)
    });
}

export default { getAPIData, base, login, signup, setAPIKey, getPreferences, setPreferences };