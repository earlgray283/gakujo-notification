import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View } from 'react-native';
import { registerRootComponent } from 'expo';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import HomeScreen from './screens/Home';
import LoginScreen from './screens/Login';
import RegisterScreen from './screens/Register';
import React, { createContext, useEffect, useState } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { Assignment } from './apis/assignment';
import AssignmentDetail from './screens/AssignmentDetail';

export type RootStackParamList = {
  Home: undefined;
  Login: undefined;
  Register: undefined;
  AssignmentDetail: { assignment: Assignment };
};

const Stack = createNativeStackNavigator<RootStackParamList>();
export const jwtTokenDispatchContext = createContext<
  React.Dispatch<React.SetStateAction<string | null | undefined>>
>(() => undefined);

export default function App() {
  const [jwtToken, setJwtToken] = useState<string | null | undefined>(
    undefined
  );
  useEffect(() => {
    (async () => {
      try {
        const jwtToken = await AsyncStorage.getItem('jwtToken');
        setJwtToken(jwtToken);
      } catch (e) {
        await AsyncStorage.clear();
      }
    })();
  }, []);
  if (jwtToken === undefined) {
    return (
      <View>
        <Text>loading...</Text>
      </View>
    );
  }

  return (
    <jwtTokenDispatchContext.Provider value={setJwtToken}>
      <NavigationContainer>
        <Stack.Navigator>
          {jwtToken ? (
            <>
              <Stack.Screen
                name='Home'
                component={HomeScreen}
                options={() => ({
                  title: '課題一覧',
                })}
              />
              <Stack.Screen
                name='AssignmentDetail'
                component={AssignmentDetail}
                options={({ route }) => ({
                  title: route.params.assignment.subjectName,
                })}
              />
            </>
          ) : (
            <>
              <Stack.Screen name='Login' component={LoginScreen} />
              <Stack.Screen name='Register' component={RegisterScreen} />
            </>
          )}
        </Stack.Navigator>
      </NavigationContainer>
    </jwtTokenDispatchContext.Provider>
  );
}

registerRootComponent(App);
