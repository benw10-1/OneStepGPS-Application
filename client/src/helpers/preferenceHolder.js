import { Requests } from ".";

function PrefHolder () {
    let data = {};
    let updateFuncs = [];

    function onUpdate(callback) {
        updateFuncs.push(callback);
        return updateFuncs.length - 1;
    }

    function removeUpdate(callback) {
        if (typeof callback === 'number') {
            updateFuncs.splice(callback, 1);
        }
        else {
            updateFuncs = updateFuncs.filter(func => func !== callback);
        }
    }

    function get() {
        return data;
    }

    async function set(data_) {
        if (data_ instanceof Promise) {
            data_ = await data_;
        }
        else if (typeof data_ === 'function') {
            data_ = await data_();
        }
        else data = data_;
        
        updateFuncs.forEach(func => func(data));

        await Requests.setPreferences(data);

        return data;
    }

    set(Requests.getAPIData());

    return {
        get,
        set,
        onUpdate,
        removeUpdate,
    }
}

export default PrefHolder();