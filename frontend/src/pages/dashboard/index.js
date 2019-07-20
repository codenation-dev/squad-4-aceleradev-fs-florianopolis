import React, { useEffect } from "react";
import Sidemenu from '../../components/sidemenu'
import "./style.css"
// import {getCustomerById} from '../../services/customerService'

const Dashboard = () => {    
    
    useEffect(() => {
        console.log("registro")
        //console.log(getCustomerById(98));
    });

    return (
        <>
            <Sidemenu />
            <div className="content">
                <h1>ASD</h1>
            </div>
        </>
    )
}

export default Dashboard;
  