class DataModel {
  final int id;
  final String nama_barang;
  final int harga;
  final int stok;

  DataModel(
      {required this.id,
      required this.nama_barang,
      required this.harga,
      required this.stok});

  // Factory method untuk parsing JSON
  factory DataModel.fromJson(Map<String, dynamic> json) {
    return DataModel(
      id: json['id'],
      nama_barang: json['nama_barang'],
      harga: json['harga'],
      stok: json['stok'],
    );
  }
}
