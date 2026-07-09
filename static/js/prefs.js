(function() {
    // Automatically update language and generation preferences in the cookie
    const pathParts = window.location.pathname.split('/').filter(p => p);
    if (pathParts.length >= 2) {
        const lang = pathParts[0];
        const gen = pathParts[1];
        document.cookie = "pokegrep-prefs=" + lang + "/" + gen + "; path=/; max-age=31536000; SameSite=Lax";
    }
})();
