import React, { useEffect } from "react";
import Sidemenu from '../../components/sidemenu'
import "./style.css"

const Dashboard = () => {    
    
    useEffect(() => {
        console.log("registro")
        //console.log(getCustomerById(98));
        //login("arthur_dent@dont_panic.com", "123")
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
  