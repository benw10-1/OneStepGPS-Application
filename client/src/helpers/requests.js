import Auth from "./auth";
// import Dummy from '../dummy/data.json';
const HOST_URL = "/";
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

    return fetch(url, options).then(res => res.json()).then(data => {
        if (data["error"]) {
            return Promise.reject(data["error"]);
        }
        return data["result"];
    }).catch(err => {
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

function convertUnitObj(obj, to='mi') {
    if (!obj) return obj;
    const { value, unit } = obj;

    if (unit === 'km/h' || unit === 'km') {
        return Math.round(value/1.60934);
    }
    if (unit === "m") {
        if (to === 'ft' ) return Math.round(value * 3.28084);
        return Math.round(value / 1609.34);
    }

    return value;
}

async function getDevices() {
    const key = Auth.getProfile()?.APIKey;
    if (!key) return Promise.reject('No API key');

    const url = HOST_URL + 'api/devices';

    return base(url).then(data => {
        if (!data) return Promise.reject('No data');
        if (data.error === "No API key") {
            setAPIKey('');
            Auth.logout();
            return Promise.reject('No data');
        }
        const devices = []
        // device schema
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
                odometer: convertUnitObj(device.latest_device_point?.device_state?.odometer ?? device.odometer),
                speed: convertUnitObj(device.latest_device_point?.device_point_detail?.speed ?? device.speed),
                altitude: convertUnitObj(device.latest_device_point?.device_point_detail?.altitude ?? device.altitude, "ft"),
                voltage: device.latest_device_point?.device_point_detail?.external_volt ?? device.external_volt,
            }
            devices.push(formatted);
        }
        // console.log(devices);
        return devices;
    })
}

async function setAPIKey(key) {
    const url = HOST_URL + 'api/setAPIKey';

    return base(url, {
        method: 'POST',
        body: JSON.stringify({
            key
        })
    }).then(data => {
        if (data.error) {
            return {
                token: "temp",
                error: data.error
            }
        }
        else if (data?.Token) {
            Auth.login(data.Token);
        }
        return data;
    })
}

async function getPreferences() {
    const key = Auth.getProfile()?.APIKey;
    if (!key) return Promise.reject('No API key');
    const url = HOST_URL + 'api/preferences';

    return base(url);
}

async function setPreferences(prefs) {
    const key = Auth.getProfile()?.APIKey;
    if (!key) return Promise.reject('No API key');
    const url = HOST_URL + 'api/preferences';

    return base(url, {
        method: 'POST',
        body: JSON.stringify(prefs)
    });
}

async function reverseGeocode(lat, lng) {
    const url = HOST_URL + 'api/reverseGeocode';

    return base(url, {
        method: 'POST',
        body: JSON.stringify({
            lat: String(lat),
            lon: String(lng)
        })
    }).then(data => {
        if (!data) return Promise.reject('No data');
        return {
            lat: data.lat,
            lng: data.lng,
            address: data.formatted_address
        };
    })
}

export default { getDevices, base, login, signup, setAPIKey, getPreferences, setPreferences, reverseGeocode };