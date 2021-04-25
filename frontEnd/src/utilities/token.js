import moment from "moment";

export function isTokenValid(token) {
    if (!token) {
        return false
    }

    let parsedToken = JSON.parse(atob(token.split('.')[1]));

    if (moment.now()/1000 > parsedToken.Expiration) {
        console.log("token expired");
        sessionStorage.removeItem("access_token");
        return false
    }

    return true
}
