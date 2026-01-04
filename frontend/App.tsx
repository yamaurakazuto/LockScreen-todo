import React, {useState} from 'react';
import {
  SafeAreaView,
  StyleSheet,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from 'react-native';

const initialTodos = [
  {id: '1', title: 'ウィジェット用データ同期を確認', completed: false},
  {id: '2', title: 'バックエンドと同期', completed: true},
];

export default function App() {
  const [title, setTitle] = useState('');
  const [todos, setTodos] = useState(initialTodos);

  const addTodo = () => {
    if (!title.trim()) {
      return;
    }
    setTodos([
      {id: Date.now().toString(), title: title.trim(), completed: false},
      ...todos,
    ]);
    setTitle('');
  };

  const toggleTodo = (id: string) => {
    setTodos((current) =>
      current.map((todo) =>
        todo.id === id ? {...todo, completed: !todo.completed} : todo,
      ),
    );
  };

  return (
    <SafeAreaView style={styles.container}>
      <Text style={styles.title}>Lock Screen Todo</Text>
      <View style={styles.inputRow}>
        <TextInput
          value={title}
          onChangeText={setTitle}
          placeholder="新しいTodoを入力"
          style={styles.input}
        />
        <TouchableOpacity onPress={addTodo} style={styles.addButton}>
          <Text style={styles.addButtonText}>追加</Text>
        </TouchableOpacity>
      </View>
      <View style={styles.list}>
        {todos.map((todo) => (
          <TouchableOpacity
            key={todo.id}
            style={styles.todoItem}
            onPress={() => toggleTodo(todo.id)}>
            <Text style={[styles.todoText, todo.completed && styles.completed]}>
              {todo.title}
            </Text>
          </TouchableOpacity>
        ))}
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 24,
    backgroundColor: '#0f172a',
  },
  title: {
    fontSize: 24,
    fontWeight: '600',
    color: '#f8fafc',
    marginBottom: 16,
  },
  inputRow: {
    flexDirection: 'row',
    gap: 12,
    marginBottom: 20,
  },
  input: {
    flex: 1,
    backgroundColor: '#1e293b',
    color: '#f8fafc',
    borderRadius: 12,
    paddingHorizontal: 12,
    paddingVertical: 10,
  },
  addButton: {
    backgroundColor: '#38bdf8',
    borderRadius: 12,
    paddingHorizontal: 16,
    justifyContent: 'center',
  },
  addButtonText: {
    color: '#0f172a',
    fontWeight: '600',
  },
  list: {
    gap: 12,
  },
  todoItem: {
    padding: 14,
    borderRadius: 12,
    backgroundColor: '#1e293b',
  },
  todoText: {
    color: '#e2e8f0',
    fontSize: 16,
  },
  completed: {
    textDecorationLine: 'line-through',
    color: '#94a3b8',
  },
});
