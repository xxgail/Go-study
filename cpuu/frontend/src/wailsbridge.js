export default {
    // The main function
    // Passes the main Wails object to the callback if given.
    Start: function (callback) {
        if (callback) {
            window.wails.Events.On("wails:ready", callback);
        }
    }
};