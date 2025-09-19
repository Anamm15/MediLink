import { Paperclip, Send } from "lucide-react";

// Contoh data chat
const mockMessages = [
  {
    id: 1,
    sender: "doctor",
    text: "Selamat pagi, Pak Budi. Ada yang bisa saya bantu?",
  },
  {
    id: 2,
    sender: "patient",
    text: "Pagi, Dok. Saya ada keluhan ruam di kulit lengan saya.",
  },
  { id: 3, sender: "patient", type: "image", url: "/ruam.jpg" }, // Placeholder gambar
  {
    id: 4,
    sender: "doctor",
    text: "Baik, terima kasih fotonya. Sejak kapan ruam ini muncul?",
  },
  {
    id: 5,
    sender: "doctor",
    type: "prescription",
    items: ["Hydrocortisone 1% Cream", "Cetirizine 10mg Tablet"],
  },
];

export const ChatInterface = () => {
  return (
    <div className="flex-grow flex flex-col bg-white h-full">
      {/* Message Area */}
      <div className="flex-grow p-6 space-y-6 overflow-y-auto">
        {mockMessages.map((msg) => (
          <div
            key={msg.id}
            className={`flex items-end gap-3 ${
              msg.sender === "patient" ? "justify-end" : "justify-start"
            }`}
          >
            {msg.sender === "doctor" && (
              <div className="w-8 h-8 rounded-full bg-slate-200 flex-shrink-0"></div>
            )}
            <div
              className={`max-w-md p-4 rounded-2xl ${
                msg.sender === "patient"
                  ? "bg-cyan-500 text-white rounded-br-none"
                  : "bg-slate-100 text-gray-800 rounded-bl-none"
              }`}
            >
              {msg.type === "prescription" ? (
                <div>
                  <h4 className="font-bold mb-2">Resep Digital Diterbitkan</h4>
                  <ul className="list-disc list-inside text-sm mb-3">
                    {msg.items?.map((item) => (
                      <li key={item}>{item}</li>
                    ))}
                  </ul>
                  <button className="w-full text-xs font-semibold bg-white text-cyan-600 py-2 rounded-md">
                    Tebus Resep
                  </button>
                </div>
              ) : msg.type === "image" ? (
                <img
                  src={msg.url}
                  alt="Uploaded content"
                  className="rounded-lg max-w-xs"
                />
              ) : (
                <p>{msg.text}</p>
              )}
            </div>
          </div>
        ))}
      </div>
      {/* Input Area */}
      <div className="bg-white border-t border-gray-200 p-4">
        <div className="flex items-center gap-3 bg-slate-100 rounded-lg p-2">
          <button className="p-2 text-gray-500 hover:text-cyan-500">
            <Paperclip className="w-5 h-5" />
          </button>
          <textarea
            placeholder="Ketik pesan Anda..."
            rows={1}
            className="w-full bg-transparent resize-none focus:outline-none text-sm"
          ></textarea>
          <button className="p-2 bg-cyan-500 text-white rounded-md hover:bg-cyan-600">
            <Send className="w-5 h-5" />
          </button>
        </div>
      </div>
    </div>
  );
};
