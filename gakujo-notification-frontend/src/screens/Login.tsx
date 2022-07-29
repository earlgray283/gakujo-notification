import AsyncStorage from '@react-native-async-storage/async-storage';
import { NavigationProp, useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { fetchAssignments } from '../apis/assignment';
import { RootStackParamList } from 'App';
import { Controller, useForm } from 'react-hook-form';
import { View, Text, TextInput, Button, StyleSheet } from 'react-native';
import { signin } from '../apis/auth';
import React, { useState } from 'react';

interface LoginInput {
  username: string;
  password: string;
}

function LoginScreen(): JSX.Element {
  const [errorMessage, setErrorMessage] = useState<string | undefined>();
  const { handleSubmit, control } = useForm<LoginInput>();
  const navigation =
    useNavigation<NativeStackNavigationProp<RootStackParamList, 'Home'>>();

  return (
    <View>
      {errorMessage && <Text style={styles.error}>{errorMessage}</Text>}

      <Controller
        name='username'
        control={control}
        defaultValue=''
        render={({ field: { onBlur, onChange, value } }) => (
          <TextInput
            style={styles.input}
            onBlur={onBlur}
            onChangeText={onChange}
            value={value}
            textContentType='username'
          />
        )}
      />

      <Controller
        name='password'
        control={control}
        defaultValue=''
        render={({ field: { onBlur, onChange, value } }) => (
          <TextInput
            style={styles.input}
            onBlur={onBlur}
            onChangeText={onChange}
            value={value}
            secureTextEntry
            textContentType='newPassword'
          />
        )}
      />

      <Button
        title='register'
        onPress={handleSubmit(async (data) => {
          try {
            navigation.navigate('Register');
          } catch (e) {
            if (e instanceof Error) {
              setErrorMessage(e.message);
            }
          }
        })}
      />

      <Button
        title='Login'
        onPress={handleSubmit(async (data) => {
          try {
            const jwtToken = await signin(data.username, data.password);
            await AsyncStorage.setItem('jwtToken', jwtToken);
          } catch (e) {
            if (e instanceof Error) {
              setErrorMessage(e.message);
            }
          }
        })}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  input: {
    height: 40,
    margin: 12,
    borderWidth: 1,
    padding: 10,
  },
  error: {
    backgroundColor: '#f8d7da',
    margin: 12,
    padding: 10,
    borderRadius: 10,
    overflow: 'hidden',
  },
  horizontal: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
  },
});

export default LoginScreen;
