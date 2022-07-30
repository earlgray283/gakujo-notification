import AsyncStorage from '@react-native-async-storage/async-storage';
import { NavigationProp, useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { fetchAssignments } from '../apis/assignment';
import { jwtTokenDispatchContext, RootStackParamList } from '../App';
import { Controller, useForm } from 'react-hook-form';
import { View, Text, TextInput, Button, StyleSheet } from 'react-native';
import { signin, signup } from '../apis/auth';
import React, { useContext, useState } from 'react';

interface RegisterInput {
  username: string;
  password: string;
  gakujoId: string;
  gakujoPassword: string;
}

function RegisterScreen(): JSX.Element {
  const [errorMessage, setErrorMessage] = useState<string | undefined>();
  const {
    handleSubmit,
    control,
    formState: { errors },
  } = useForm<RegisterInput>();
  const setJwtToken = useContext(jwtTokenDispatchContext);

  return (
    <View>
      {errorMessage && <Text style={styles.error}>{errorMessage}</Text>}
      {errors.gakujoId && errors.gakujoId.message && (
        <Text style={styles.error}>{errors.gakujoId.message}</Text>
      )}
      {errors.gakujoPassword && errors.gakujoPassword.message && (
        <Text style={styles.error}>{errors.gakujoPassword.message}</Text>
      )}
      {errors.username && errors.username.message && (
        <Text style={styles.error}>{errors.username.message}</Text>
      )}
      {errors.password && errors.password.message && (
        <Text style={styles.error}>{errors.password.message}</Text>
      )}

      <Controller
        name='username'
        control={control}
        defaultValue=''
        rules={{ required: 'username must be set' }}
        render={({ field: { onBlur, onChange, value } }) => (
          <TextInput
            style={styles.input}
            onBlur={onBlur}
            onChangeText={onChange}
            value={value}
            textContentType='username'
            placeholder='username'
          />
        )}
      />

      <Controller
        name='password'
        control={control}
        defaultValue=''
        rules={{ required: 'password must be set' }}
        render={({ field: { onBlur, onChange, value } }) => (
          <TextInput
            style={styles.input}
            onBlur={onBlur}
            onChangeText={onChange}
            value={value}
            secureTextEntry
            textContentType='newPassword'
            placeholder='password'
          />
        )}
      />

      <Controller
        name='gakujoId'
        control={control}
        defaultValue=''
        rules={{ required: 'gakujoId must be set' }}
        render={({ field: { onBlur, onChange, value } }) => (
          <TextInput
            style={styles.input}
            onBlur={onBlur}
            onChangeText={onChange}
            value={value}
            textContentType='username'
            placeholder='gakujo id'
          />
        )}
      />

      <Controller
        name='gakujoPassword'
        control={control}
        defaultValue=''
        rules={{ required: 'gakujo password must be set' }}
        render={({ field: { onBlur, onChange, value } }) => (
          <TextInput
            style={styles.input}
            onBlur={onBlur}
            onChangeText={onChange}
            value={value}
            secureTextEntry
            textContentType='newPassword'
            placeholder='gakujo password'
          />
        )}
      />

      <Button
        title='register'
        onPress={handleSubmit(async (data) => {
          try {
            const jwtToken = await signup(
              data.username,
              data.password,
              data.gakujoId,
              data.gakujoPassword
            );
            await AsyncStorage.setItem('jwtToken', jwtToken);
            setJwtToken(jwtToken);
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

export default RegisterScreen;
