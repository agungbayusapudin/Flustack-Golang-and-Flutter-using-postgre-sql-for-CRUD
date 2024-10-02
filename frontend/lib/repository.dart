import 'dart:convert';

import 'package:frontend/data.dart';
import 'package:http/http.dart' as http;

class Repository {
  final _baseurl = 'http://localhost:8080/api/get';

  Future getData() async {
    final response = await http.get(Uri.parse(_baseurl));

    if (response.statusCode == 200) {
      final Map<String, dynamic> jsonData = jsonDecode(response.body);

      List<DataModel> items = (jsonData['data'] as List)
          .map((item) => DataModel.fromJson(Map<String, dynamic>.from(item)))
          .toList();

      return items;
    }
  }

  Future updateData(int id, String nama_barang, int harga, int stok) async {
    final _baseurl = 'http://localhost:8080/api/update/$id';

    final response = await http.put(Uri.parse(_baseurl),
        headers: <String, String>{
          'Content-Type': 'aplication/json; charset=UTF-8'
        },
        body: jsonEncode(<String, dynamic>{
          'nama_barang': nama_barang,
          'harga': harga,
          'stok': stok
        }));
    if (response.statusCode != 200) {
      throw Exception('errrr');
    }
  }

  Future uploadData(String nama_barang, int harga, int stok) async {
    final _baseUrl = 'http://localhost:8080/api/tambah';

    final response = await http.post(Uri.parse(_baseUrl),
        headers: <String, String>{
          'Content-Type': 'aplication/json; charset=UTF-8'
        },
        body: jsonEncode(<String, dynamic>{
          'nama_barang': nama_barang,
          'harga': harga,
          'stok': stok
        }));
    if (response.statusCode != 200) {
      throw Exception('errrr');
    }
  }

  Future deleteData(int id) async {
    final _baseUrl = 'http://localhost:8080/api/delete/$id';

    final response = await http.delete(
      Uri.parse(_baseUrl),
      headers: <String, String>{
        'Content-Type': 'aplication/json; charset=UTF-8'
      },
    );

    if (response.statusCode != 200) {
      throw Exception('errrr');
    }
  }
}
