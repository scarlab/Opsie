import { Actions, useCsDispatch, useCsSelector } from "@/cs-redux";
import UserSlice from "@/cs-redux/slices/user.slice";
import { ArrowLeft } from "lucide-react";
import { useEffect } from "react";
import { useSearchParams } from "react-router-dom";
import NewUser from "./NewUser";
import UserList from "./UserList";
import UserDetails from "./UserDetails";
import Tabs from "./Tabs";

export default function UsersView() {
    const dispatch = useCsDispatch();
    const [searchParams, setSearchParams] = useSearchParams();
    const userId = searchParams.get("user");
    const { user } = useCsSelector(state => state.user);

    useEffect(() => {
        if (userId) {
            dispatch(Actions.user.GetUserById({ id: userId! }));
            dispatch(Actions.team.GetAllTeamsOfUser({ user_id: userId! }));
        }
    }, [userId])

    return (
        <div className="space-y-5 pb-20">
            <div className="flex items-center justify-between ">
                <h1 className="text-2xl font-bold flex items-center gap-3">
                    {userId && <ArrowLeft
                        onClick={() => {
                            searchParams.delete("user");
                            setSearchParams(searchParams);
                            dispatch(UserSlice.actions.removeUser({}));
                        }} />
                    }

                    {userId ? user?.display_name : "Users"}
                </h1>
                <NewUser />
            </div>


            {userId ?
                <>
                    <UserDetails />
                    <Tabs />
                </>
                : <UserList />
            }

        </div>
    )
}
