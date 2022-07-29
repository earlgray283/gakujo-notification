import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import React, { useEffect, useState } from 'react';
import {
  View,
  Text,
  FlatList,
  ActivityIndicator,
  StyleSheet,
  Button,
  TouchableOpacity,
} from 'react-native';
import { Assignment, fetchAssignments } from '../apis/assignment';
import { RootStackParamList } from '../App';
import { toSimpleDateString } from '../libs/date';

function HomeScreen(): JSX.Element {
  const navigation =
    useNavigation<
      NativeStackNavigationProp<RootStackParamList, 'AssignmentDetail'>
    >();
  const [astmts, setAstmts] = useState<Assignment[] | null | undefined>(
    undefined
  );
  useEffect(() => {
    try {
      (async () => {
        const astmts = await fetchAssignments();
        setAstmts(astmts);
      })();
    } catch (e) {
      if (e instanceof Error) {
        setAstmts([]);
      }
    }
  }, []);

  if (!astmts) {
    return <ActivityIndicator size='large' />;
  }

  return (
    <View>
      <FlatList
        data={astmts.filter((astmt) => astmt.status === '受付中')}
        renderItem={({ item }) => (
          <TouchableOpacity
            style={styles.item}
            onPress={() => {
              navigation.navigate('AssignmentDetail', { assignment: item });
            }}
          >
            <Text>{item.subjectName}</Text>
            <Text>{toSimpleDateString(item.until)}</Text>
          </TouchableOpacity>
        )}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  item: {
    padding: 12,
    backgroundColor: '#99b7dc',
    marginLeft: 15,
    marginRight: 15,
    marginTop: 5,
    borderRadius: 10,
  },
});

export default HomeScreen;
