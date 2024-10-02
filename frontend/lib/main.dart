import 'dart:convert'; // Untuk decode data JSON
import 'dart:ffi';
import 'package:flutter/material.dart';
import 'package:frontend/data.dart';
import 'package:frontend/repository.dart';
import 'package:http/http.dart' as http;

// Model class

void main() {
  runApp(MaterialApp(
    home: HomeScreen(),
  ));
}

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  bool isLoading = true;
  List<DataModel> barang = [];
  Repository repository = Repository();

  getData() async {
    setState(() {
      isLoading = true;
    });
    try {
      barang = await repository.getData();
      setState(() {
        isLoading = false;
      }); // Update UI after data is fetched
    } catch (error) {
      print('Error fetching data: $error');
    }
  }

  updateData(int id, String nama_barang, int harga, int stok) async {
    try {
      await repository.updateData(id, nama_barang, harga, stok);
      await getData();
    } catch (e) {
      throw Exception(e);
    }
  }

  uploadData(String nama_barang, int harga, int stok) async {
    try {
      await repository.uploadData(nama_barang, harga, stok);
      await getData();
    } catch (e) {
      print('errr: $e');
    }
  }

  deleteData(int id) async {
    try {
      await repository.deleteData(id);
      await getData();
    } catch (e) {
      throw Exception(e);
    }
  }

  @override
  void initState() {
    getData();
    super.initState();
  }

  void showItemDialog({int? id, String? nama_barang, int? harga, int? stok}) {
    TextEditingController nama_barangController =
        TextEditingController(text: nama_barang);
    TextEditingController hargaController =
        TextEditingController(text: harga?.toString() ?? '');
    TextEditingController stokController =
        TextEditingController(text: stok?.toString() ?? '');

    showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: Text(id == null ? 'Create item' : 'update Item'),
            content: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                TextField(
                  controller: nama_barangController,
                  decoration: InputDecoration(labelText: 'Nama Barang'),
                ),
                TextField(
                  controller: hargaController,
                  decoration: InputDecoration(labelText: 'harga'),
                  keyboardType: TextInputType.number,
                ),
                TextField(
                  controller: stokController,
                  decoration: InputDecoration(labelText: 'stok'),
                  keyboardType: TextInputType.number,
                ),
              ],
            ),
            actions: [
              TextButton(
                  onPressed: () {
                    Navigator.of(context).pop();
                  },
                  child: Text('Cancel')),
              TextButton(
                  onPressed: () {
                    try {
                      if (id == null) {
                        uploadData(
                          nama_barangController.text,
                          int.parse(hargaController.text),
                          int.parse(stokController.text), // Correct argument
                        );
                      } else {
                        updateData(
                          id,
                          nama_barangController.text,
                          int.parse(hargaController.text),
                          int.parse(stokController.text),
                        );
                      }
                      Navigator.of(context).pop(); // Close dialog after action
                    } catch (e) {
                      // Handle parsing errors
                      print('Error: $e');
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(content: Text('Please enter valid data.')),
                      );
                    }
                  },
                  child: Text(id == null ? 'create' : 'Update'))
            ],
          );
        });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("menampilkan data dari API"),
      ),
      body: isLoading
          ? Center(
              child: CircularProgressIndicator(),
            )
          : ListView.builder(
              itemCount: barang.length,
              itemBuilder: (context, index) {
                var item = barang[index];

                return ListTile(
                  title: Text(item.nama_barang),
                  subtitle: Text(item.harga.toString()),
                  trailing: Row(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      IconButton(
                          onPressed: () {
                            showItemDialog(
                                id: item.id,
                                nama_barang: item.nama_barang,
                                harga: item.harga,
                                stok: item.stok);
                          },
                          icon: Icon(Icons.edit)),
                      IconButton(
                          onPressed: () {
                            deleteData(item.id);
                          },
                          icon: Icon(Icons.delete))
                    ],
                  ),
                );
              }),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          showItemDialog();
        },
        child: Icon(Icons.add),
      ),
    );
  }
}
