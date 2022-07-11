import Auth from "./auth";
// import Dummy from '../dummy/data.json';
const HOST_URL = 'http://localho.st:3000/';
// fetch wrapper
async function base(url, override={}) {
    let token = Auth.getToken();

    const options = {
        method: 'GET',
        headers: {
            'Content-Type': '*',
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Headers": "*",
            "Access-Control-Allow-Methods": "*",
            'cache-control': 'no-cache',
            "Authorization": `Bearer ${token ?? ''}`
        },
        mode: 'cors',
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

async function getDevices() {
    const key = Auth.getProfile().APIKey;
    if (!key) return Promise.reject('No API key');

    const url = HOST_URL + 'api/getDevices';

    return base(url).then(data => {
        const devices = []
        for (const device of data?.result_list ?? data ?? []) {
            const formatted = {
                device_id: device.device_id,
                display_name: device.display_name,
                online: device.online,
                heading: device.heading ?? device.latest_device_point.angle,
                lat: device.latest_device_point?.lat ?? device.lat,
                lng: device.latest_device_point?.lng ?? device.lng,
                drive_status: device.latest_device_point?.device_state?.drive_status ?? device.drive_status,
                drive_status_begin_time: device.latest_device_point?.device_state?.drive_status_begin_time ?? device.drive_status_begin_time,
            }
            devices.push(formatted);
        }
        // console.log(devices);
        return devices;
    });
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

export default { getDevices, base, login, signup, setAPIKey, getPreferences, setPreferences };