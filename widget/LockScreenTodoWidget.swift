import WidgetKit
import SwiftUI

struct TodoItem: Identifiable, Codable {
    let id: String
    let title: String
    let completed: Bool
}

struct TodoEntry: TimelineEntry {
    let date: Date
    let todos: [TodoItem]
}

struct TodoProvider: TimelineProvider {
    func placeholder(in context: Context) -> TodoEntry {
        TodoEntry(date: Date(), todos: [TodoItem(id: "1", title: "ロック画面で確認", completed: false)])
    }

    func getSnapshot(in context: Context, completion: @escaping (TodoEntry) -> Void) {
        completion(loadEntry())
    }

    func getTimeline(in context: Context, completion: @escaping (Timeline<TodoEntry>) -> Void) {
        let entry = loadEntry()
        let timeline = Timeline(entries: [entry], policy: .after(Date().addingTimeInterval(900)))
        completion(timeline)
    }

    private func loadEntry() -> TodoEntry {
        let defaults = UserDefaults(suiteName: "group.com.example.lockscreentodo")
        let data = defaults?.data(forKey: "todos")
        let todos: [TodoItem]
        if let data, let decoded = try? JSONDecoder().decode([TodoItem].self, from: data) {
            todos = decoded
        } else {
            todos = [TodoItem(id: "1", title: "ウィジェットデータなし", completed: false)]
        }
        return TodoEntry(date: Date(), todos: Array(todos.prefix(3)))
    }
}

struct LockScreenTodoWidgetEntryView: View {
    var entry: TodoProvider.Entry

    var body: some View {
        VStack(alignment: .leading, spacing: 6) {
            ForEach(entry.todos) { todo in
                HStack(alignment: .center, spacing: 6) {
                    Circle()
                        .fill(todo.completed ? Color.green : Color.blue)
                        .frame(width: 6, height: 6)
                    Text(todo.title)
                        .font(.caption)
                        .lineLimit(1)
                }
            }
        }
        .padding(12)
    }
}

@main
struct LockScreenTodoWidget: Widget {
    let kind: String = "LockScreenTodoWidget"

    var body: some WidgetConfiguration {
        StaticConfiguration(kind: kind, provider: TodoProvider()) { entry in
            LockScreenTodoWidgetEntryView(entry: entry)
        }
        .configurationDisplayName("Lock Screen Todo")
        .description("最新のTodoをロック画面に表示します。")
        .supportedFamilies([.systemSmall, .systemMedium])
    }
}
