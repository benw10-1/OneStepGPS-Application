import Requests from './requests';

function locationCache() {
    let cache = {}

    async function requestLocation(latLonArr) {
        return new Promise((resolve) => {
            const [lat, lon] = latLonArr;

            if (cache[lat] && cache[lat][lon]) {
                resolve(cache[lat][lon]);
                return;
            }

            Requests.reverseGeocode(lat, lon).then(data => {
                if (!data) return;
                if (!cache[lat]) cache[lat] = {};

                cache[lat][lon] = data.address;

                resolve(data.address);
            }).catch(err => {
                console.log(err);
                resolve(null);
            })
        })
    }

    return {
        requestLocation
    }
}

export default locationCache();