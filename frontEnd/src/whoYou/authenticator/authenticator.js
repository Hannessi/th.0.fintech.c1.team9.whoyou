import config from 'react-global-configuration'
import {jsonRpcRequestRaw} from '../helper'

class Authenticator {
    static Login ({username, password}) {
        return jsonRpcRequestRaw({
            url: config.get('golang'),
            method: 'Authenticator.Login',
            request: {
                username,
                password
            },
        })
    }
}

export default Authenticator
