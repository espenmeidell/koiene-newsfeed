import React from "react"
import ReactDOM from "react-dom"
import "./index.css"
import App from "./App"
import registerServiceWorker from "./registerServiceWorker"
import { initializeIcons } from "office-ui-fabric-react/lib/Icons"

initializeIcons()

ReactDOM.render(<App />, document.getElementById("root"))
registerServiceWorker()
