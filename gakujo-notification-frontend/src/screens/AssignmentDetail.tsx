import { useNavigation } from '@react-navigation/native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { Card } from '@rneui/base';
import React, { useEffect, useState } from 'react';
import {
  View,
  Text,
  FlatList,
  ActivityIndicator,
  StyleSheet,
  ScrollView,
} from 'react-native';
import { Assignment, fetchAssignments } from '../apis/assignment';
import { RootStackParamList } from '../App';
import { toSimpleDateString } from '../libs/date';

type Props = NativeStackScreenProps<RootStackParamList, 'AssignmentDetail'>;

function AssignmentDetailScreen(props: Props): JSX.Element {
  const asmt = props.route.params.assignment;
  return (
    <ScrollView>
      <Card>
        <Card.Title>
          {asmt.title} {`(締切: ${toSimpleDateString(asmt.until)})`}
        </Card.Title>
        <Card.Divider />
        <Text>{asmt.description}</Text>
      </Card>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  card: {
    padding: 12,
    backgroundColor: '#99b7dc',
    marginLeft: 15,
    marginRight: 15,
    marginTop: 5,
    borderRadius: 10,
  },
});

export default AssignmentDetailScreen;
