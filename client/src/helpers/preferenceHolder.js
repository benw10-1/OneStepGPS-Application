import { Requests } from ".";

function PrefHolder () {
    let data = {
        isPrefs__: true
    };
    let updateFuncs = [];

    function onUpdate(callback) {
        if (typeof callback !== 'function') return -1
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

    async function set(data_, func_=null) {
        if (data_ instanceof Promise) {
            data_ = await data_;
        }
        else if (typeof data_ === 'function') {
            data_ = await data_(data);
        }

        data = { ...data, ...data_ };
        
        updateFuncs.forEach(func => {if (func_ !== func) func(data);});
        
        await Requests.setPreferences(data);

        return data;
    }

    async function preferenceReq() {
        return Requests.getPreferences().then(data_ => {
            if (!data_) return
            data = { ...data, ...data_ };
            updateFuncs.forEach(func => {func(data)});
        });
    }

    return {
        get,
        set,
        onUpdate,
        removeUpdate,
        preferenceReq
    }
}

export default PrefHolder();