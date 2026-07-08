(function() {
    const cookieName = "pokegrep-prefs";

    function getCookie(name) {
        const value = "; " + document.cookie;
        const parts = value.split("; " + name + "=");
        if (parts.length === 2) return parts.pop().split(";").shift();
        return null;
    }

    function setCookie(name, value, days) {
        let expires = "";
        if (days) {
            const date = new Date();
            date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
            expires = "; expires=" + date.toUTCString();
        }
        document.cookie = name + "=" + (value || "")  + expires + "; path=/";
    }

    let pref = getCookie(cookieName);

    if (!pref) {
        const browserLang = (navigator.language || navigator.userLanguage || "en").toLowerCase();
        const lang = browserLang.startsWith("fr") ? "fr" : "en";
        const gen = "generation-i";

        pref = lang + "/" + gen;
        setCookie(cookieName, pref, 365);
    }

    window.location.href = "/" + pref + "/";
})();
