import RootLoader from '@/components/loader/RootLoader';
import { Actions, useCsDispatch, useCsSelector } from '@/cs-redux';
import AuthSlice from '@/cs-redux/slices/auth.slice';
import { getLocalAuthUser, removeLocalAuthUser } from '@/helpers/auth.helper';
import { useEffect, useState } from 'react';
import { Outlet, useLocation, useNavigate } from 'react-router-dom';

export default function RootLayout() {
    const dispatch = useCsDispatch();
    const navigate = useNavigate();
    const location = useLocation();

    const { authUser, loading } = useCsSelector((state) => state.auth);
    const [sessionInitializing, setSessionInitializing] = useState(true);
    const [onboardingRequired, setOnboardingRequired] = useState(false);

    useEffect(() => {
        const init = async () => {
            try {
                // Check if system is onboarded (owner exists)
                const { payload: { count } } = await dispatch(Actions.user.GetOwnerCount());

                if (count <= 0) {
                    removeLocalAuthUser();
                    setOnboardingRequired(true); // mark onboarding
                    navigate('/onboarding', { replace: true });
                    return; // skip rest
                }

                // Try restore cached user
                const cachedUser = getLocalAuthUser();
                if (cachedUser && cachedUser.id && cachedUser.is_active) {
                    dispatch(AuthSlice.actions.restoreAuthUser(cachedUser));
                } else {
                    removeLocalAuthUser();
                    // Verify session server-side
                    await dispatch(Actions.auth.Session());
                }
            } catch (err) {
                console.error("RootLayout init error:", err);
            } finally {
                setSessionInitializing(false);
            }
        };

        init();
    }, [dispatch, navigate]);

    // Redirect logic after session fully initialized
    useEffect(() => {
        if (!sessionInitializing && !loading && !onboardingRequired) {
            if (!authUser && !location.pathname.startsWith('/auth')) {
                navigate('/auth/login', { replace: true });
            } else if (authUser && location.pathname.startsWith('/auth')) {
                navigate('/', { replace: true });
            }
        }
    }, [sessionInitializing, loading, authUser, onboardingRequired, navigate, location.pathname]);

    // Block render until session + onboarding check is complete
    if (sessionInitializing || loading) {
        return <RootLoader />;
    }

    // If onboarding required, Outlet will never render (redirect happened)
    return <Outlet />;
}
