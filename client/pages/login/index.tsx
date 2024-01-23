import React from "react";

const index = () => {
    return (
        <div className="flex items-center justify-center min-w-full min-h-screen">
            <form className="flex flex-col md:w-1/5" action="">
                <input
                    placeholder="email"
                    className="p-3 mt-8 rounded-md border-border-grey focus:outline-none focus:outline-blue"
                />
                <input
                    placeholder="password"
                    className="p-3 mt-4 rounded-md border-border-grey focus:outline-none focus:outline-blue"
                    type="passowrd"
                />
                <button>login</button>
            </form>
        </div>
    )
}

export default index