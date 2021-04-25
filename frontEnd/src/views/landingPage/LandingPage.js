import React from 'react'
import {
    Button,
    Typography,
    TextField,
    withStyles,
} from '@material-ui/core'

const styles = () => ({
    hover: {
        '&:hover': {
            backgroundColor: '#434334',
        },
        borderRadius: '10px',
        cursor: 'pointer',
    }
});

const pageStates = {
    landingPage: 0,
    facialScan: 1,
    enterPhoneNumber: 2,
}

class LandingPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            pageState: pageStates.landingPage,

        }

        this.handleMovePastFacial = this.handleMovePastFacial.bind(this)
        this.handleRegisterNewUser = this.handleRegisterNewUser.bind(this)
    }

    componentDidMount() {
    }

    render() {
        return (
            <div>
                <div style={{
                    position: 'absolute',
                    height: '99vh',
                    width: '100%',
                    zIndex: '10',
                    backgroundColor: '#eee',
                }}>
                    <div style={{
                        top: '50%',
                        left: '50%',
                        transform: 'translate(-50%,-50%)',
                        position: 'relative',
                        backgroundColor: 'white',
                        width: '414px',
                        height: '736px'
                    }}>
                        {this.renderPage()}
                    </div>
                </div>
            </div>
        )
    }

    renderPage() {
        const {
            pageState
        } = this.state

        switch (pageState) {
            case pageStates.landingPage:
                return this.renderLandingPage()
            case pageStates.facialScan:
                return this.renderFacialScan()
            case pageStates.enterPhoneNumber:
                return this.renderEnterPhoneNumber()
            default:
                return this.renderLandingPage()
        }
    }


    renderLandingPage() {
        return (
            <div
                style={{
                    display: "grid",
                    gridTemplateRows: "1fr auto auto auto 1fr",
                    gridRowGap: "10px",
                    alignContent: 'center',
                    height: "100%",
                    padding: "0 60px"
                }}
            >
                <div/>
                <div style={{
                    textAlign: 'center'
                }}>
                    <Typography style={{fontSize: '210%'}}>
                        IDaaS
                    </Typography>
                </div>
                <Button variant="contained" color="primary" onClick={this.handleRegisterNewUser}>
                    Register New User
                </Button>
                <Button variant="contained" color="primary">
                    Find User
                </Button>
            </div>
        )
    }

    renderFacialScan() {
        return (
            <div
                style={{
                    display: "grid",
                    gridTemplateRows: "1fr auto auto auto 1fr",
                    gridRowGap: "10px",
                    alignContent: 'center',
                    height: "100%",
                    padding: "0 60px",
                    backgroundColor: "#f7f7f7",
                }}
            >
                <div/>
                <div style={{
                    textAlign: 'center'
                }}>
                    <img src={require("../../images/app/scan.gif")} alt={"scan"} width={300}/>
                </div>
                <Button variant="contained" color="primary" onClick={this.handleMovePastFacial}>
                    Next
                </Button>
            </div>
        )
    }

    renderEnterPhoneNumber() {
        return (
            <div
                style={{
                    display: "grid",
                    gridTemplateRows: "1fr auto auto auto 1fr",
                    gridRowGap: "10px",
                    alignContent: 'center',
                    height: "100%",
                    padding: "0 60px",
                    backgroundColor: "#f7f7f7",
                }}
            >
                <div/>
                <div style={{
                    textAlign: 'center'
                }}>
                    <Typography>Enter phone number to enable MFA</Typography>
                </div>
                <div style={{
                    padding: "40px 0px"
                }}>
                    <TextField
                        fullWidth
                        value={"432093"}
                        style={{
                            textAlign: "center"
                        }}
                        variant="outlined"
                    />
                </div>
                <Button variant="contained" color="primary" onClick={this.handleMovePastFacial}>
                    Next
                </Button>
            </div>
        )
    }


    handleRegisterNewUser() {
        this.setState({pageState: pageStates.facialScan})
    }

    handleMovePastFacial() {
        this.setState({pageState: pageStates.enterPhoneNumber})
    }
}

export default withStyles(styles)(LandingPage)
