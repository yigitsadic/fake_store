import React from "react";
import {User} from "./user";

export interface AuthInterface {
    currentUser?: User | undefined;
    setCurrentUser?:  React.Dispatch<React.SetStateAction<User | undefined>>
}
